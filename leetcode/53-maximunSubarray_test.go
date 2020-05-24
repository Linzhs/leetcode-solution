package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, v := range tests {
		got := maxSubArrayByDivideConquer(v.nums)
		if got != v.want {
			t.Errorf("test case (%v) want %v but got %v", v, v.want, got)
		}
	}
}
