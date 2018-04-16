package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		num             string
		expected_result int
	}{
		{"I", 1},
		{"MDCLXVI", 1666},
		{"XXXIX", 39},
		{"MMMCMXCIX", 3999},
	}

	for i, ta := range tt {
		actual_result := romanToInt(ta.num)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
