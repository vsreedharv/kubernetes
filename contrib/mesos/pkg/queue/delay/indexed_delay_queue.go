/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package delay

import (
	"container/heap"
	"fmt"
	"reflect"
	"time"

	"github.com/pivotal-golang/clock"

	"k8s.io/kubernetes/contrib/mesos/pkg/queue"
	"k8s.io/kubernetes/contrib/mesos/pkg/queue/priority"
	"k8s.io/kubernetes/pkg/util/sets"
)

// IndexedDelayQueue is a thread-safe, time-based queue with uniquely identified items.
//
// Adding or offering an item with an ID matching an item still in the queue will
// invoke the configured SchedulingPolicy and supplied ValueReplacementPolicy to
// determine the desired behavior.
//
// Items can only be popped after their event time has been reached.
// Use Add to specify a delay duration and have the event time calculated.
// Use Offer to specify a specific event time.
type IndexedDelayQueue struct {
	*Queue
	// items mapped by UniqueID - item set must match queue contents
	items            map[string]priority.Item
	schedulingPolicy SchedulingPolicy
	clock            clock.Clock
}

func NewIndexedDelayQueue(clock clock.Clock) *IndexedDelayQueue {
	return &IndexedDelayQueue{
		Queue: NewDelayQueue(clock),
		items: map[string]priority.Item{},
		clock: clock,
	}
}

// Add inserts an uniquely identified item into the queue,
// provided another item with the same ID is not already present.
func (q *IndexedDelayQueue) Add(d UniqueDelayed, rp ValueReplacementPolicy) {
	q.push(&indexedAddedItem{
		Item: priority.NewItem(d, NewDelayedPriority(d, q.clock)),
		id:   d.GetUID(),
	}, rp)
}

// indexedAddedItem wraps an item that was added to the queue.
type indexedAddedItem struct {
	priority.Item
	id string
}

func (di indexedAddedItem) GetUID() string {
	return di.id
}

// Push puts the item back into the queue (with the same ID)
func (di *indexedAddedItem) Push(queue heap.Interface) {
	dq := queue.(*IndexedDelayQueue)
	dq.push(di, KeepExisting)
}

func (q *IndexedDelayQueue) Offer(d UniqueScheduled, rp ValueReplacementPolicy) bool {
	if p, ok := NewScheduledPriority(d); ok {
		q.push(&indexedOfferedItem{
			Item: priority.NewItem(d, p),
			id:   d.GetUID(),
		}, rp)
		return true
	}
	return false
}

func (q *IndexedDelayQueue) push(newItem UniqueItem, rp ValueReplacementPolicy) {
	q.lock.Lock()
	defer q.lock.Unlock()
	id := newItem.GetUID()
	item, exists := q.items[id]
	if !exists {
		item = newItem
		heap.Push(q.Queue, item)
		q.items[id] = item
	} else {
		// replace existing item
		newValue := rp.Value(item.Value(), newItem.Value())
		oldPriority := item.Priority().(Priority)
		newPriority := newItem.Priority().(Priority)
		newPriority = q.schedulingPolicy.EventTime(oldPriority, newPriority)
		item = priority.NewItem(newValue, newPriority)

		switch i := newItem.(type) {
		case *indexedAddedItem:
			item = &indexedAddedItem{Item: item, id: i.id}
		case *indexedOfferedItem:
			item = &indexedOfferedItem{Item: item, id: i.id}
		default:
			panic(fmt.Sprintf("unsupported newItem type: %v", reflect.TypeOf(i)))
		}

		heap.Fix(q.Queue, item.Index())
		q.items[id] = item
	}
	q.cond.Broadcast()
}

// indexedOfferedItem represents an item that was offered to the queue.
type indexedOfferedItem struct {
	priority.Item
	id string
}

func (oi indexedOfferedItem) GetUID() string {
	return oi.id
}

// Push offers the item to be rescheduled (with a new ID)
func (oi *indexedOfferedItem) Push(queue heap.Interface) {
	dq := queue.(*IndexedDelayQueue)
	dq.Offer(oi.Value().(UniqueScheduled), KeepExisting)
}

// Delete removes an item. It doesn't add it to the queue, because
// this implementation assumes the consumer only cares about the objects,
// not their priority order.
func (q *IndexedDelayQueue) Delete(id string) {
	q.lock.Lock()
	defer q.lock.Unlock()
	delete(q.items, id)
}

// List returns a list of all the items.
func (q *IndexedDelayQueue) List() []queue.UniqueID {
	q.lock.RLock()
	defer q.lock.RUnlock()
	list := make([]queue.UniqueID, 0, len(q.items))
	for _, item := range q.items {
		list = append(list, item.Value().(UniqueDelayed))
	}
	return list
}

// ContainedIDs returns a stringset.StringSet containing all IDs of the stored items.
// This is a snapshot of a moment in time, and one should keep in mind that
// other go routines can add or remove items after you call this.
func (q *IndexedDelayQueue) ContainedIDs() sets.String {
	q.lock.RLock()
	defer q.lock.RUnlock()
	set := sets.String{}
	for id := range q.items {
		set.Insert(id)
	}
	return set
}

// Get returns the requested item, or sets exists=false.
func (q *IndexedDelayQueue) Get(id string) (queue.UniqueID, bool) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if item, exists := q.items[id]; exists {
		return item.Value().(queue.UniqueID), true
	}
	return nil, false
}

// Variant of DelayQueue.Pop() for UniqueDelayed items
func (q *IndexedDelayQueue) Await(timeout time.Duration) queue.UniqueID {
	cancel := make(chan struct{})
	ch := make(chan interface{}, 1)
	go func() { ch <- q.pop(cancel) }()
	var x interface{}
	timer := q.clock.NewTimer(timeout)
	select {
	case <-timer.C():
		close(cancel)
		x = <-ch
	case x = <-ch:
		timer.Stop()
	}
	if x != nil {
		return x.(queue.UniqueID)
	}
	return nil
}

// Variant of DelayQueue.Pop() for UniqueDelayed items
func (q *IndexedDelayQueue) Pop() interface{} {
	return q.pop(nil).(queue.UniqueID)
}

// variant of DelayQueue.Pop that implements optional cancellation
func (q *IndexedDelayQueue) pop(cancel chan struct{}) interface{} {
	next := func() pushableItem {
		q.lock.Lock()
		defer q.lock.Unlock()
		for {
			for q.Len() == 0 {
				signal := make(chan struct{})
				go func() {
					defer close(signal)
					q.cond.Wait()
				}()
				select {
				case <-cancel:
					// we may not have the lock yet, so
					// broadcast to abort Wait, then
					// return after lock re-acquisition
					q.cond.Broadcast()
					<-signal
					return nil
				case <-signal:
					// we have the lock, re-check
					// the queue for data...
				}
			}
			//TODO: should this just be q.Queue? If so, it deadlocks...
			x := heap.Pop(q.Queue.Queue)
			item := x.(pushableItem)
			unique := item.Value().(queue.UniqueID)
			uid := unique.GetUID()
			if _, ok := q.items[uid]; !ok {
				// item was deleted, keep looking
				continue
			}
			delete(q.items, uid)
			return item
		}
	}
	return q.Queue.pop(next, cancel)
}
