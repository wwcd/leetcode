package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		nums             []int
		target           int
		expected_indices []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for i, ta := range tt {
		actual_indices := twoSum(ta.nums, ta.target)
		if !reflect.DeepEqual(ta.expected_indices, actual_indices) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_indices, actual_indices)
		}
	}
}
