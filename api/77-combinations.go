package api

func combine(n int, k int) [][]int {

	if n <= 0 || k <= 0 || k > n {
		return nil
	}

	var result [][]int
	backtrackingCombine(&result, []int{}, n, k, 1)
	return result
}

func backtrackingCombine(result *[][]int, curSlice []int, n, k, pos int) {

	if len(curSlice) == k {
		*result = append(*result, curSlice)
		return
	}

	for i := pos; i-1+(k-len(curSlice)) <= n; i++ {
		s := make([]int, 0, len(curSlice))
		s = append(s, curSlice...)
		s = append(s, i)
		backtrackingCombine(result, s, n, k, i+1)
	}
}
