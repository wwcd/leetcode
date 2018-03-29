package main

import (
	"reflect"
	"testing"
)

func TestLongtestPalindrome(t *testing.T) {
	tt := []struct {
		str             string
		rows            int
		expected_result string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"ABC", 2, "ACB"},
	}

	for i, ta := range tt {
		actual_result := convert(ta.str, ta.rows)
		if !reflect.DeepEqual(ta.expected_result, actual_result) {
			t.Errorf("[FAILED] caseid: %d, expected: %v, actual: %v", i, ta.expected_result, actual_result)
		}
	}
}
