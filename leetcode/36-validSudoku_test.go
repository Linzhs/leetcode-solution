package leetcode

import "testing"

func TestValidSudoku(t *testing.T) {
	tests := []struct {
		board [][]byte
		want  bool
	}{
		{[][]byte{
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}, false},
	}

	for _, v := range tests {
		got := isValidSudoku(v.board)
		if got != v.want {
			t.Errorf("test case (%+v) got %v but want %v", v, got, v.want)
		}
	}
}
