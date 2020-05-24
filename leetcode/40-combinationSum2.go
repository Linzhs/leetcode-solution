package leetcode

import "sort"

// 1 2 2 2 5
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 || target == 0 {
		return nil
	}

	sort.Sort(sort.IntSlice(candidates))

	var result [][]int
	backtraceSum2(&result, candidates, []int{}, 0, 0, target)

	return result
}

func backtraceSum2(result *[][]int, candidates []int, sumSet []int, curPos, curSum, target int) {
	switch {
	case curSum > target:
	case curSum == target:
		s := make([]int, 0, len(sumSet))
		s = append(s, sumSet...)
		*result = append(*result, s)
	default:
		for i := curPos; i < len(candidates); i++ {
			if i > curPos && candidates[i] == candidates[i-1] {
				continue
			}
			s := make([]int, 0, len(sumSet)+1)
			s = append(s, sumSet...)
			s = append(s, candidates[i])
			backtraceSum2(result, candidates, s, i+1, curSum+candidates[i], target)
		}
	}
}
