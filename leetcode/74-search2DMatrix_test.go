package leetcode

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13, false},
		{[][]int{{1, 1}}, 2, false},
	}

	for _, v := range tests {
		got := searchMatrix(v.matrix, v.target)
		if got != v.want {
			t.Errorf("test case (%+v) want %v but got %v", v, v.want, got)
		}
	}
}
