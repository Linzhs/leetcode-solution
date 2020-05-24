package leetcode

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend, divisor int
		want              int
	}{
		{15, 3, 5},
		{15, 2, 7},
		{10, 3, 3},
		{7, -3, -2},
	}

	for _, v := range tests {
		got := divide(v.dividend, v.divisor)
		if got != v.want {
			t.Errorf("test case %+v want %d but got %d", v, v.want, got)
		}
	}
}
