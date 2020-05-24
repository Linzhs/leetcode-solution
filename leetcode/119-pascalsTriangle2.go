package leetcode

func getRow(rowIndex int) (result []int) {
	for i := 0; i <= rowIndex; i++ {
		result = append(result, 1)
		for j := i - 1; j >= 1; j-- {
			result[j] = result[j-1] + result[j]
		}
	}

	return
}
