package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		str             string
		expected_result int
	}{
		{"123", 123},
		{"-123", -123},
		{"   123", 123},
		{"   -123", -123},
		{"   123   ", 123},
		{"   -123   ", -123},
		{"+-2", 0},
		{"    010", 10},
		{"-2147483648", -2147483648},
		{"b2147483648", 0},
	}

	for i, ta := range tt {
		actual_result := myAtoi(ta.str)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
