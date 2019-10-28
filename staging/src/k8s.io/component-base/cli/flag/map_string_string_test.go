/*
Copyright 2017 The Kubernetes Authors.

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

package flag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMapStringString(t *testing.T) {
	var nilMap map[string]string
	cases := []struct {
		desc   string
		m      *MapStringString
		expect string
	}{
		{"nil", NewMapStringString(&nilMap), ""},
		{"empty", NewMapStringString(&map[string]string{}), ""},
		{"one key", NewMapStringString(&map[string]string{"one": "foo"}), "one=foo"},
		{"two keys", NewMapStringString(&map[string]string{"one": "foo", "two": "bar"}), "one=foo,two=bar"},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			assert.Equalf(t, c.expect, c.m.String(), "expect %q but got %q", c.expect, c.m.String())
		})
	}
}

func TestSetMapStringString(t *testing.T) {
	var nilMap map[string]string
	cases := []struct {
		desc           string
		vals           []string
		start          *MapStringString
		expect         *MapStringString
		err            string
		expectedToFail bool
	}{
		// we initialize the map with a default key that should be cleared by Set
		{
			"clears defaults",
			[]string{""},
			NewMapStringString(&map[string]string{"default": ""}),
			&MapStringString{
				initialized: true,
				Map:         &map[string]string{},
				NoSplit:     false,
			},
			"",
			false,
		},
		// make sure we still allocate for "initialized" maps where Map was initially set to a nil map
		{
			"allocates map if currently nil",
			[]string{""},
			&MapStringString{initialized: true, Map: &nilMap},
			&MapStringString{
				initialized: true,
				Map:         &map[string]string{},
				NoSplit:     false,
			},
			"",
			false,
		},
		// for most cases, we just reuse nilMap, which should be allocated by Set, and is reset before each test case
		{
			"empty",
			[]string{""},
			NewMapStringString(&nilMap),
			&MapStringString{
				initialized: true,
				Map:         &map[string]string{},
				NoSplit:     false,
			},
			"",
			false,
		},
		{
			"one key",
			[]string{"one=foo"},
			NewMapStringString(&nilMap),
			&MapStringString{
				initialized: true,
				Map:         &map[string]string{"one": "foo"},
				NoSplit:     false,
			},
			"",
			false,
		},
		{
			"two keys",
			[]string{"one=foo,two=bar"},
			NewMapStringString(&nilMap),
			&MapStringString{
				initialized: true,
				Map:         &map[string]string{"one": "foo", "two": "bar"},
				NoSplit:     false,
			},
			"",
			false,
		},
		{
			"one key, multi flag invocation only",
			[]string{"one=foo,bar"},
			NewMapStringStringNoSplit(&nilMap),
			&MapStringString{
				initialized: true,
				Map:         &map[string]string{"one": "foo,bar"},
				NoSplit:     true,
			},
			"",
			false,
		},
		{
			"two keys, multi flag invocation only",
			[]string{"one=foo,bar", "two=foo,bar"},
			NewMapStringStringNoSplit(&nilMap),
			&MapStringString{
				initialized: true,
				Map:         &map[string]string{"one": "foo,bar", "two": "foo,bar"},
				NoSplit:     true,
			},
			"",
			false,
		},
		{
			"two keys, multiple Set invocations",
			[]string{"one=foo", "two=bar"},
			NewMapStringString(&nilMap),
			&MapStringString{
				initialized: true,
				Map:         &map[string]string{"one": "foo", "two": "bar"},
				NoSplit:     false,
			},
			"",
			false,
		},
		{
			"two keys with space",
			[]string{"one=foo, two=bar"},
			NewMapStringString(&nilMap),
			&MapStringString{
				initialized: true,
				Map:         &map[string]string{"one": "foo", "two": "bar"},
				NoSplit:     false,
			},
			"",
			false,
		},
		{
			"empty key",
			[]string{"=foo"},
			NewMapStringString(&nilMap),
			&MapStringString{
				initialized: true,
				Map:         &map[string]string{"": "foo"},
				NoSplit:     false,
			},
			"",
			false,
		},
		{
			"missing value",
			[]string{"one"},
			NewMapStringString(&nilMap),
			nil,
			"malformed pair, expect string=string",
			true,
		},
		{
			"no target",
			[]string{"a:foo"},
			NewMapStringString(nil),
			nil,
			"no target (nil pointer to map[string]string)",
			true,
		},
	}
	for _, c := range cases {
		nilMap = nil
		t.Run(c.desc, func(t *testing.T) {
			var err error
			for _, val := range c.vals {
				if err = c.start.Set(val); err != nil {
					break
				}
			}
			if c.expectedToFail {
				assert.Equalf(t, c.err, err.Error(), "expected error %s but got %v", c.err, err)
			} else {
				assert.Nil(t, err, "unexpected error: %v", err)
				assert.Equalf(t, c.expect, c.start, "expected %#v but got %#v", c.expect, c.start)
			}
		})
	}
}

func TestEmptyMapStringString(t *testing.T) {
	var nilMap map[string]string
	cases := []struct {
		desc   string
		val    *MapStringString
		expect bool
	}{
		{"nil", NewMapStringString(&nilMap), true},
		{"empty", NewMapStringString(&map[string]string{}), true},
		{"populated", NewMapStringString(&map[string]string{"foo": ""}), false},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			result := c.val.Empty()
			assert.Equalf(t, c.expect, result, "expect %t but got %t", c.expect, result)
		})
	}
}
