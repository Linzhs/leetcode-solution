package api

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {

	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{7}, {2, 2, 3}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}

	for _, test := range tests {
		got := combinationSum(test.candidates, test.target)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("test case (candidates: %v) want %v but got %v", test.candidates, test.want, got)
		}
	}
}
