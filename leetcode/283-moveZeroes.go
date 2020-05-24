package leetcode

// [0,1,0,3,12]
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	lastNonZeroIndex := 0
	for _, v := range nums {
		if v != 0 {
			nums[lastNonZeroIndex] = v
			lastNonZeroIndex++
		}
	}

	for i := lastNonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}
