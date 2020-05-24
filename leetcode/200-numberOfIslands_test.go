package leetcode

import "testing"

func TestNumIsLands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{grid: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, want: 1},
	}

	for _, v := range tests {
		got := numIslands(v.grid)
		if got != v.want {
			t.Errorf("test case (%+v) want %v but got %v", v, v.want, got)
		}
	}
}
