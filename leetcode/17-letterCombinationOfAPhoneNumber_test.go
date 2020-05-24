package leetcode

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		digits string
		want   []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for _, test := range tests {
		got := letterCombinations(test.digits)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("test case (digits: %s) got %#+v but want %#+v", test.digits, got, test.want)
		}
	}
}
