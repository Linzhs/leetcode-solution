package leetcode

// nums为升序、无重复元素数组
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	for lo, hi := 0, len(nums)-1; lo <= hi; {
		mid := (lo + hi) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[lo] > nums[mid] {
			hi = mid - 1
			continue
		}
		lo = mid + 1
	}

	return nums[0]
}
