package leetcode

func generate(numRows int) (result [][]int) {
	if numRows <= 0 {
		return
	}

	result = append(result, []int{1})

	for i := 1; i < numRows; i++ {
		var s []int
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				s = append(s, 1)
				continue
			}

			s = append(s, result[i-1][j-1]+result[i-1][j])

		}

		result = append(result, s)
	}

	return
}
