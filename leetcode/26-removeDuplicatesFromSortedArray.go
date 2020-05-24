package leetcode

func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
