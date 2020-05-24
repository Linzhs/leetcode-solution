package leetcode

//[[1, 1, 1, 1, 1, 1, 1],
// [1, 2, 2, 2, 2, 2, 1],
// [1, 2, 3, 3, 3, 2, 1],
// [1, 2, 2, 2, 2, 2, 1],
// [1, 1, 1, 1, 1, 1, 1]]
//
// 1 2 3
// 4 5 6
// 7 8 9
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	var res []int
	if len(matrix) == 1 {
		res = append(res, matrix[0]...)
		return res
	}

	c0, cn := 0, len(matrix[0])-1
	r0, rn := 0, len(matrix)-1

	for c0 <= cn && r0 <= rn {
		for i := c0; i <= cn; i++ {
			res = append(res, matrix[r0][i])
		}
		for i := r0 + 1; i <= rn; i++ {
			res = append(res, matrix[i][cn])
		}

		if c0 < cn && r0 < rn {
			for i := cn - 1; i > c0; i-- {
				res = append(res, matrix[rn][i])
			}
			for i := rn; i > r0; i-- {
				res = append(res, matrix[i][c0])
			}
		}

		c0++
		cn--
		r0++
		rn--
	}

	return res
}
