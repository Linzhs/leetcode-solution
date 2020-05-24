package leetcode

import "sort"

func subsetsWithDup(nums []int) (result [][]int) {
	result = append(result, []int{})
	if len(nums) == 0 {
		return
	}

	sort.Sort(sort.IntSlice(nums))

	for i := 1; i <= len(nums); i++ {
		backtrackingSubSetsWithDup(&result, []int{}, nums, 0, i)
	}

	return
}

func backtrackingSubSetsWithDup(result *[][]int, curNums, nums []int, pos, curLen int) {
	if len(curNums) == curLen {
		s := make([]int, 0, len(curNums))
		s = append(s, curNums...)
		*result = append(*result, s)
	}

	for i := pos; i < len(nums); i++ {
		if i > pos && nums[i] == nums[i-1] {
			continue
		}

		s := make([]int, 0, len(curNums)+1)
		s = append(s, curNums...)
		s = append(s, nums[i])
		backtrackingSubSetsWithDup(result, s, nums, i+1, curLen)
	}

}
