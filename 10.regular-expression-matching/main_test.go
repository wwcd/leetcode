package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		s               string
		p               string
		expected_result bool
	}{
		{"aa", "a", false},
		{"aa", "aa", true},
		{"aaa", "aa", false},
		{"a", "aa", false},
		{"aa", "a*", true},
		{"ba", "a*", false},
		{"aa", ".*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"c", ".*c", true},
		{"abc", ".*c", true},
		{"ab", ".*c", false},
		{"", "c*c*", true},
	}

	for i, ta := range tt {
		actual_result := isMatch(ta.s, ta.p)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
