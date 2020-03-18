package api

// 交替的单调
func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) / 2
		switch nums[mid] < nums[mid+1] {
		case true:
			lo = mid + 1
		case false:
			hi = mid
		}
	}

	return lo
}
