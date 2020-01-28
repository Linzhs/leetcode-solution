package api

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{nums: []int{1, 0, -1, 0, -2, 2}, target: 0, want: [][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}}},
		{nums: []int{-1, 0, 1, 2, -1, -4}, target: -1, want: [][]int{{-4, 0, 1, 2}, {-1, -1, 0, 1}}},
		{nums: []int{-3, -2, -1, 0, 0, 1, 2, 3}, target: 0, want: [][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
	}

	for _, v := range tests {
		got := fourSum(v.nums, v.target)
		if !cmp.Equal(got, v.want) {
			t.Errorf("test case (nums:%v, target:%d) want %v but got %v", v.nums, v.target, v.want, got)
		}
	}
}
