package leetcode

import "testing"

func TestRotatedSearch2(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   bool
	}{
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{1, 1}, 0, false},
		{[]int{8, 9, 2, 3, 4}, 0, true},
		{[]int{1, 3, 1, 1, 1}, 3, true},
	}

	for _, v := range tests {
		got := search2(v.nums, v.target)
		if got != v.want {
			t.Errorf("test case (%+v) got %v but want %v", v, got, v.want)
		}
	}
}
