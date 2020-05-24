package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	leftPositionIndex := extremeBinarySearch(nums, target, true)
	if leftPositionIndex == len(nums) || nums[leftPositionIndex] != target {
		return []int{-1, -1}
	}

	return []int{leftPositionIndex, extremeBinarySearch(nums, target, false) - 1}
}

func extremeBinarySearch(nums []int, target int, isLeft bool) int {
	lo, hi := 0, len(nums) // 招不到时为len(nums)
	for lo < hi {
		mid := (lo + hi) / 2
		switch {
		case nums[mid] > target, isLeft && nums[mid] == target:
			hi = mid
		default:
			lo = mid + 1
		}
	}

	return lo
}
