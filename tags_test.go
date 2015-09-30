// Copyright 2015 Vic Demuzere
//
// Use of this source code is governed by the MIT license.

package irc

import (
	"reflect"
	"testing"
)

var tagsTests = [...]*struct {
	parsed Tags
	raw    string
}{
	{
		parsed: Tags{
			"aaa":             "bbb",
			"ccc":             "",
			"example.com/ddd": "eee",
		},
		raw: "aaa=bbb;ccc;example.com/ddd=eee",
	},
}

// Unable to test this for now due to the random order in maps.
//func TestTags_String(t *testing.T) {
//	var s string
//
//	for i, test := range tagsTests {
//
//		// Convert the taglist
//		s = test.parsed.String()
//
//		// Result should be the same as the value in raw.
//		if s != test.raw {
//			t.Errorf("Failed to stringify taglist %d:", i)
//			t.Logf("Output: %s", s)
//			t.Logf("Expected: %s", test.raw)
//		}
//	}
//}

func TestTags_Len(t *testing.T) {
	var l int

	for i, test := range tagsTests {

		l = test.parsed.Len()

		// Result should be the same as the value in raw.
		if l != len(test.raw) {
			t.Errorf("Failed to calculate taglist length %d:", i)
			t.Logf("Output: %d", l)
			t.Logf("Expected: %d", len(test.raw))
		}
	}
}

func TestParseTags(t *testing.T) {
	var ta Tags

	for i, test := range tagsTests {

		// Parse the prefix
		ta = ParseTags(test.raw)

		// Result Tags should be the same as the value in parsed.
		if !reflect.DeepEqual(ta, test.parsed) {
			t.Errorf("Failed to parse taglist %d:", i)
			t.Logf("Output: %#v", ta)
			t.Logf("Expected: %#v", test.parsed)
		}
	}
}
