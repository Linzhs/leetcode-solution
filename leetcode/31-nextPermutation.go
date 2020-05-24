package leetcode

// [1,5,8,4,7,6,5,3,1]
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	index := len(nums) - 2

	// <<==
	for index >= 0 && nums[index] >= nums[index+1] {
		index--
	}

	// ==>
	if index >= 0 {
		exchangeIndex := len(nums) - 1
		for exchangeIndex >= 0 && nums[exchangeIndex] <= nums[index] {
			exchangeIndex--
		}
		nums[exchangeIndex], nums[index] = nums[index], nums[exchangeIndex]
	}

	// reverse
	for i, j := index+1, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
