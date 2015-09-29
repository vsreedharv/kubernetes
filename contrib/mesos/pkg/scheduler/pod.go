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

package scheduler

import (
	"fmt"
	"time"

	"k8s.io/kubernetes/contrib/mesos/pkg/queue/delay"
	"k8s.io/kubernetes/contrib/mesos/pkg/queue/historical"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/cache"
)

// wrapper for the k8s pod type so that we can define additional methods on a "pod"
type Pod struct {
	*api.Pod
	eventTime *time.Time
	delay     *time.Duration
	notify    delay.BreakChan
}

// implements Copyable
func (p *Pod) Copy() historical.Copyable {
	if p == nil {
		return nil
	}
	//TODO(jdef) we may need a better "deep-copy" implementation
	pod := *(p.Pod)
	return &Pod{Pod: &pod}
}

// implements queue.UniqueID
func (p *Pod) GetUID() string {
	if id, err := cache.MetaNamespaceKeyFunc(p.Pod); err != nil {
		panic(fmt.Sprintf("failed to determine pod id for '%+v'", p.Pod))
	} else {
		return id
	}
}

// implements queue/delay.Scheduled
func (dp *Pod) EventTime() (time.Time, bool) {
	if dp.eventTime != nil {
		return *(dp.eventTime), true
	}
	return time.Time{}, false
}

// implements queue/delay.Delayed
func (dp *Pod) GetDelay() time.Duration {
	if dp.delay != nil {
		return *(dp.delay)
	}
	return 0
}

// implements queue/delay.Breakout
func (p *Pod) Breaker() delay.BreakChan {
	return p.notify
}

func (p *Pod) String() string {
	displayDeadline := "<none>"
	if eventTime, ok := p.EventTime(); ok {
		displayDeadline = eventTime.String()
	}
	return fmt.Sprintf("{pod:%v, eventTime:%v, delay:%v}", p.Pod.Name, displayDeadline, p.GetDelay())
}
