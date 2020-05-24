package leetcode

// len(nums) > 1
func productExceptSelf(nums []int) (result []int) {
	result = make([]int, len(nums))
	result[0] = 1

	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * rightProduct
		rightProduct *= nums[i]
	}

	return
}
