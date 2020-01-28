package api

import "sort"

// test case:
// [1,1,2]
// [1,1,1,2]
// [0,1,0,0,9]
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	copyNums := make([]int, len(nums))
	copy(copyNums, nums)

	sort.Sort(sort.IntSlice(copyNums))

	var result [][]int
	backtrackPermuteUnique(&result, copyNums, 0)

	return result
}

func backtrackPermuteUnique(result *[][]int, curPermutation []int, start int) {
	if start == len(curPermutation) {
		s := make([]int, 0, len(curPermutation))
		s = append(s, curPermutation...)
		*result = append(*result, s)
	}

	set := make(map[int]bool, len(curPermutation))

	for i := start; i < len(curPermutation); i++ {
		// pruning
		//if i > start && curPermutation[i] == curPermutation[i-1] {
		//	continue
		//}
		if set[curPermutation[i]] {
			continue
		}
		set[curPermutation[i]] = true

		// swap
		curPermutation[i], curPermutation[start] = curPermutation[start], curPermutation[i]

		// backtracking
		s := make([]int, 0, len(curPermutation))
		s = append(s, curPermutation...)
		backtrackPermuteUnique(result, s, start+1)

		// reduction
		curPermutation[i], curPermutation[start] = curPermutation[start], curPermutation[i]
	}
}
