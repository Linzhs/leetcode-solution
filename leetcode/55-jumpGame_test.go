package leetcode

import "testing"

func TestJumpGame(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}

	for _, v := range tests {
		got := canJump(v.nums)
		if got != v.want {
			t.Errorf("test case (%v) want %v but got %v", v, v.want, got)
		}
	}
}
