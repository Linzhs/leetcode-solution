package leetcode

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB", false},
	}

	for _, v := range tests {
		got := exist(v.board, v.word)
		if got != v.want {
			t.Errorf("test case (%+v) want %v but got %v", v, v.want, got)
		}
	}
}
