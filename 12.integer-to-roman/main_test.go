package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		num             int
		expected_result string
	}{
		{1, "I"},
		{1666, "MDCLXVI"},
		{39, "XXXIX"},
		{3999, "MMMCMXCIX"},
	}

	for i, ta := range tt {
		actual_result := intToRoman(ta.num)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
