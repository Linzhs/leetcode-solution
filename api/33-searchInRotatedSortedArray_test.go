package api

import "testing"

func TestRotatedSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3}, 0, -1},
		{[]int{3, 1}, 3, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{3, 4, 5, 6, 1, 2}, 2, 5},
		{[]int{8, 9, 2, 3, 4}, 9, 1},
	}

	for _, v := range tests {
		got := search(v.nums, v.target)
		if got != v.want {
			t.Errorf("test case (%+v) got %v but want %v", v, got, v.want)
		}
	}
}
