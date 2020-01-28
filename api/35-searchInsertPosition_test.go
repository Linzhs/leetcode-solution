package api

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}

	for _, v := range tests {
		got := searchInsert(v.nums, v.target)
		if got != v.want {
			t.Errorf("test case (%+v) got %v but want %v", v, got, v.want)
		}
	}
}
