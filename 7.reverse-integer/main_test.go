package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		x               int
		expected_result int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	}

	for i, ta := range tt {
		actual_result := reverse(ta.x)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
