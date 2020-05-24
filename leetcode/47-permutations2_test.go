package leetcode

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{nums: []int{0, 1, 0, 0, 9}, want: [][]int{
			{0, 0, 0, 1, 9},
			{0, 0, 0, 9, 1},
			{0, 0, 1, 0, 9},
			{0, 0, 1, 9, 0},
			{0, 0, 9, 0, 1},
			{0, 0, 9, 1, 0},
			{0, 1, 0, 0, 9},
			{0, 1, 0, 9, 0},
			{0, 1, 9, 0, 0},
			{0, 9, 0, 0, 1},
			{0, 9, 0, 1, 0},
			{0, 9, 1, 0, 0},
			{1, 0, 0, 0, 9},
			{1, 0, 0, 9, 0},
			{1, 0, 9, 0, 0},
			{1, 9, 0, 0, 0},
			{9, 0, 0, 0, 1},
			{9, 0, 0, 1, 0},
			{9, 0, 1, 0, 0},
			{9, 1, 0, 0, 0},
		}},
	}

	for _, v := range tests {
		got := permuteUnique(v.nums)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("test case (%+v) got %v but want %v", v, got, v.want)
		}
	}
}
