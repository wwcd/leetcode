package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		x               int
		expected_result bool
	}{
		{123, false},
		{121, true},
		{-121, false},
	}

	for i, ta := range tt {
		actual_result := isPalindrome(ta.x)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
