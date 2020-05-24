package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	rotateIndex := findRotateIndex(nums)
	switch {
	case target == nums[rotateIndex]:
		return rotateIndex
	case rotateIndex == 0:
		return binSearch(nums, 0, len(nums)-1, target)
	case target < nums[0]:
		return binSearch(nums, rotateIndex, len(nums)-1, target)
	default:
		return binSearch(nums, 0, rotateIndex, target)
	}
}

func findRotateIndex(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return 0
	}

	for lo, hi := 0, len(nums)-1; lo <= hi; {
		mid := (lo + hi) / 2
		if nums[mid] > nums[mid+1] {
			return mid + 1
		}

		if nums[lo] > nums[mid] {
			hi = mid - 1
			continue
		}
		lo = mid + 1
	}

	return 0
}

func binSearch(nums []int, lo, hi, target int) int {
	for lo <= hi {
		mid := (lo + hi) / 2
		switch {
		case nums[mid] < target:
			lo = mid + 1
		case nums[mid] > target:
			hi = mid - 1
		default:
			return mid
		}
	}

	return -1
}
