package leetcode

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{nums: []int{2, 0, 2, 1, 1, 0}, want: []int{0, 0, 1, 1, 2, 2}},
		{nums: []int{1, 0, 2, 1, 1, 2}, want: []int{0, 1, 1, 1, 2, 2}},
	}

	for _, v := range tests {
		sortColors(v.nums)
		if !reflect.DeepEqual(v.nums, v.want) {
			t.Errorf("test case (%+v) want %v but got %v", v, v.want, v.nums)
		}
	}
}
