package leetcode

func combinationSum3(k, n int) (result [][]int) {
	combinationSumDfs(&result, []int{}, n, k, 1)
	return
}

func combinationSumDfs(result *[][]int, curSlice []int, n, k, idx int) {
	if n < 0 || (len(curSlice) == k && n != 0) {
		return
	}

	if len(curSlice) == k && n == 0 {
		s := make([]int, 0, len(curSlice))
		s = append(s, curSlice...)
		*result = append(*result, s)
		return
	}

	for i := idx; i < 10; i++ {
		s := make([]int, 0, len(curSlice)+1)
		s = append(s, curSlice...)
		s = append(s, i)
		combinationSumDfs(result, s, n-i, k, i+1)
	}
}
