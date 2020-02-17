package api

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || (len(matrix) == 1 && len(matrix[0]) == 1) {
		return
	}

	m, n := len(matrix), len(matrix[0])

	// mark
	needSweepCol := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch {
			case matrix[i][j] != 0:
				// continue loop
			case j == 0:
				needSweepCol = true
			default:
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// sweep
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

	if needSweepCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
