package api

func combinationSum(candidates []int, target int) [][]int {

	if len(candidates) == 0 || target == 0 {
		return nil
	}

	var result [][]int
	backtrackSum(&result, []int{}, candidates, target, 0)
	return result
}

func backtrackSum(result *[][]int, curList, candidates []int, target, sum int) {

	if sum > target {
		return
	}

	if sum == target {
		s := make([]int, 0, len(curList))
		s = append(s, curList...)
		*result = append(*result, s)
		return
	}

	for i, v := range candidates {
		s := make([]int, 0, len(curList)+1)
		s = append(s, curList...)
		s = append(s, v)
		backtrackSum(result, s, candidates[i:], target, sum+v)
	}
}
