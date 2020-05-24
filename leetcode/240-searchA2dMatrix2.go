package leetcode

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	row, col := len(matrix)-1, 0
	for row >= 0 && col < len(matrix[0]) {
		switch {
		case matrix[row][col] > target:
			row--
		case matrix[row][col] < target:
			col++
		default:
			return true
		}
	}

	return false
}
