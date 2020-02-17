package api

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {

	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}

	for _, test := range tests {
		got := combinationSum2(test.candidates, test.target)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("test case (candidates: %v) want %v but got %v", test.candidates, test.want, got)
		}
	}
}
