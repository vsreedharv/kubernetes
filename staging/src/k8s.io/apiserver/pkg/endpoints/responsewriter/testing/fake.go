/*
Copyright 2021 The Kubernetes Authors.

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

package testing

import (
	"bufio"
	"net"
	"net/http"
	"testing"
)

var _ http.ResponseWriter = &FakeResponseWriter{}

// FakeResponseWriter implements http.ResponseWriter,
// it is used for testing purpose only
type FakeResponseWriter struct{}

func (fw *FakeResponseWriter) Header() http.Header          { return http.Header{} }
func (fw *FakeResponseWriter) WriteHeader(code int)         {}
func (fw *FakeResponseWriter) Write(bs []byte) (int, error) { return len(bs), nil }

// For HTTP2 an http.ResponseWriter object implements
// http.Flusher and http.CloseNotifier.
// It is used for testing purpose only
type FakeResponseWriterFlusherCloseNotifier struct {
	*FakeResponseWriter
}

func (fw *FakeResponseWriterFlusherCloseNotifier) Flush()                   {}
func (fw *FakeResponseWriterFlusherCloseNotifier) CloseNotify() <-chan bool { return nil }

// For HTTP/1.x an http.ResponseWriter object implements
// http.Flusher, http.CloseNotifier and http.Hijacker.
// It is used for testing purpose only
type FakeResponseWriterFlusherCloseNotifierHijacker struct {
	*FakeResponseWriterFlusherCloseNotifier
}

func (fw *FakeResponseWriterFlusherCloseNotifierHijacker) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return nil, nil, nil
}

// AssertResponseWriterInterfaceCompatibility asserts that the given
// http.ResponseWriter objects, both inner and outer, are compatible in terms
// of implementation of the following interfaces -
//   - http.Flusher
//   - responsewriter.FlusherError
//   - http.CloseNotifier
//   - http.Hijacker (applicable to http/1x only)
//
// When a given (inner) http.ResponseWriter object is decorated, the derived
// http.ResponseWriter object (outer) should implement the same interface(s)
// as the original (inner) http.ResponseWriter
func AssertResponseWriterInterfaceCompatibility(t *testing.T, inner, outer http.ResponseWriter) {
	t.Helper()

	_, innerFlushable := inner.(http.Flusher)
	_, outerFlushable := outer.(http.Flusher)
	if innerFlushable != outerFlushable {
		t.Errorf("Expected both the inner and outer http.ResponseWriter object to be compatible with http.Flusher, but got - inner: %t, outer: %t", innerFlushable, outerFlushable)
	}

	//nolint:staticcheck // SA1019
	_, innerCloseNotifiable := inner.(http.CloseNotifier)
	//nolint:staticcheck // SA1019
	_, outerCloseNotifiable := outer.(http.CloseNotifier)
	if innerCloseNotifiable != outerCloseNotifiable {
		t.Errorf("Expected the inner and outer http.ResponseWriter object to be compatible with http.CloseNotifier, but got - inner: %t, outer: %t", innerCloseNotifiable, outerCloseNotifiable)
	}

	// http/1.x implements http.Hijacker, not http2
	_, innerHijackable := inner.(http.Hijacker)
	_, outerHijackable := outer.(http.Hijacker)
	if innerHijackable != outerHijackable {
		t.Errorf("Expected the inner and outer http.ResponseWriter object to be compatible with http.Hijacker, but got - inner: %t, outer: %t", innerHijackable, outerHijackable)
	}
}

func AssertResponseWriterImplementsExtendedInterfaces(t *testing.T, w http.ResponseWriter, req *http.Request) {
	t.Helper()

	_, flushable := w.(http.Flusher)
	if !flushable {
		t.Errorf("Expected the http.ResponseWriter object of type: %T to implement http.Flusher", w)
	}

	//nolint:staticcheck // SA1019
	_, closeNotifiable := w.(http.CloseNotifier)
	if !closeNotifiable {
		t.Errorf("Expected the http.ResponseWriter object of type: %T to implement http.CloseNotifier", w)
	}

	// only http/1.x implements http.Hijacker
	if req.Proto == "HTTP/1.1" {
		_, hijackable := w.(http.Hijacker)
		if !hijackable {
			t.Errorf("Expected the http.ResponseWriter object of type: %T to implement http.Hijacker", w)
		}
	}
}
