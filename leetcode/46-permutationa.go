package leetcode

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	copyNums := make([]int, len(nums))
	copy(copyNums, nums)

	var result [][]int
	backtrackPermute(&result, copyNums, 0)

	return result
}

func backtrackPermute(result *[][]int, curPermutation []int, curStart int) {
	if curStart == len(curPermutation) {
		s := make([]int, 0, len(curPermutation))
		s = append(s, curPermutation...)
		*result = append(*result, s)
	}

	for i := curStart; i < len(curPermutation); i++ {
		// swap
		curPermutation[i], curPermutation[curStart] = curPermutation[curStart], curPermutation[i]

		// backtracking
		s := make([]int, 0, len(curPermutation))
		s = append(s, curPermutation...)
		backtrackPermute(result, s, curStart+1)

		// reduction
		curPermutation[i], curPermutation[curStart] = curPermutation[curStart], curPermutation[i]
	}
}
