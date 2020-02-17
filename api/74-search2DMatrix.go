package api

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	matrixLen, n := len(matrix)*len(matrix[0]), len(matrix[0])
	for lo, hi := 0, matrixLen-1; lo <= hi; {
		mid := (lo + hi) / 2
		elem := matrix[mid/n][mid%n]

		if elem == target {
			return true
		}

		if elem > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return false
}
