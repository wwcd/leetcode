package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		strs            []string
		expected_result string
	}{
		{[]string{}, ""},
		{[]string{""}, ""},
		{[]string{"", "123"}, ""},
		{[]string{"123", ""}, ""},
		{[]string{"123"}, "123"},
		{[]string{"123", "456"}, ""},
		{[]string{"123", "123456"}, "123"},
		{[]string{"123456", "123456"}, "123456"},
		{[]string{"123456", "1234567", "123456"}, "123456"},
	}

	for i, ta := range tt {
		actual_result := longestCommonPrefix(ta.strs)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
