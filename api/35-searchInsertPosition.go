package api

// no duplicates in the array
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	lo, hi := 0, len(nums)
	for lo < hi {
		mid := (lo + hi) / 2
		switch {
		case nums[mid] > target:
			hi = mid - 1
		case nums[mid] < target:
			lo = mid + 1
		default:
			return mid
		}
	}

	return lo
}
