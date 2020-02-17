package api

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	wordBytes := []byte(word)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existBacktracking(board, wordBytes, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func existBacktracking(board [][]byte, word []byte, m, n, cur int) bool {
	if cur == len(word) {
		return true
	}

	if m >= len(board) || n >= len(board[0]) || m < 0 || n < 0 {
		return false
	}

	if board[m][n] != word[cur] {
		return false
	}

	// hit
	cur++

	// mark
	board[m][n] ^= 255

	is := existBacktracking(board, word, m+1, n, cur) ||
		existBacktracking(board, word, m, n+1, cur) || existBacktracking(board, word, m-1, n, cur) ||
		existBacktracking(board, word, m, n-1, cur)

	// unmark
	board[m][n] ^= 255

	return is
}
