package api

import "testing"

func TestRemoveDuplicates2(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5},
	}

	for _, v := range tests {
		got := removeDuplicates2(v.nums)
		if got != v.want {
			t.Errorf("test case (%+v) want %v but got %v", v, v.want, got)
		}
	}
}
