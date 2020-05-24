package leetcode

func rotateArray(nums []int, k int) {
	k %= len(nums)
	reverseRange(nums, 0, len(nums)-1)
	reverseRange(nums, 0, k-1)
	reverseRange(nums, k, len(nums)-1)
}

func reverseRange(nums []int, lo, hi int) {
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
}
