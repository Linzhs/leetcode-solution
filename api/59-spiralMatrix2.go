package api

// 1  2  3  4
// 12 13 14 5
// 11 16 15 6
// 10 9  8  7
func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return [][]int{{1}}
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	r0, rn := 0, n-1
	c0, cn := 0, n-1
	for cur := 1; c0 <= cn && r0 <= rn; {
		for i := c0; i <= cn; i++ {
			result[r0][i] = cur
			cur++
		}

		for i := r0 + 1; i <= rn; i++ {
			result[i][cn] = cur
			cur++
		}

		for i := cn - 1; i > c0; i-- {
			result[rn][i] = cur
			cur++
		}

		for i := rn; i > r0; i-- {
			result[i][c0] = cur
			cur++
		}

		c0++
		cn--
		r0++
		rn--
	}

	return result
}
