package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		height          []int
		expected_result int
	}{
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{1, 2, 1}, 2},
	}

	for i, ta := range tt {
		actual_result := maxArea(ta.height)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
