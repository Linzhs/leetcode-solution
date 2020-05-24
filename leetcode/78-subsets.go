package leetcode

// len从0~n递增回溯 nums元素各不相同
func subsets(nums []int) (result [][]int) {
	result = append(result, []int{})
	if len(nums) == 0 {
		return
	}

	for i := 1; i <= len(nums); i++ {
		subsetsBacktracking(&result, []int{}, nums, 0, i)
	}

	return
}

func subsetsBacktracking(result *[][]int, curNums, nums []int, pos, curLen int) {
	if len(curNums) == curLen {
		s := make([]int, 0, len(curNums))
		s = append(s, curNums...)
		*result = append(*result, s)
	}

	for i := pos; i < len(nums); i++ {
		s := make([]int, 0, len(curNums)+1)
		s = append(s, curNums...)
		s = append(s, nums[i])
		subsetsBacktracking(result, s, nums, i+1, curLen)
	}
}
