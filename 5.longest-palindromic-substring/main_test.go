package main

import (
	"reflect"
	"testing"
)

func TestLongtestPalindrome(t *testing.T) {
	tt := []struct {
		str             string
		expected_result string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"abcda", "a"},
		{"aaabaaaa", "aaabaaa"},
	}

	for i, ta := range tt {
		actual_result := longestPalindrome(ta.str)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
