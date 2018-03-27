package main

import (
	"reflect"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tt := []struct {
		nums1           []int
		nums2           []int
		expected_result float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for i, ta := range tt {
		actual_result := findMedianSortedArrays(ta.nums1, ta.nums2)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
