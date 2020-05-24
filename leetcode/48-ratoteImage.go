package leetcode

// 旋转 N x N 矩阵 = 转置 + 翻转
//
// [
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
// ]
func rotate(matrix [][]int) {
	if len(matrix) <= 1 {
		return
	}

	// transpose
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			// swap
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// flip
	for left, right := 0, len(matrix)-1; left < right; {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
		left++
		right--
	}
}
