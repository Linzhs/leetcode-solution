package leetcode

// search1: [4,5,6,7,0,1,2]
// search2: [2,5,6,0,0,1,2]
// nums[0]为中间值 主要位置
func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return nums[0] == target
	}

	rotatePos := findRotatePosition(nums)
	switch {
	case nums[rotatePos] == target:
		return true
	case rotatePos == 0: // 未旋转
		return binarySearch2(nums, 0, len(nums)-1, target)
	case nums[0] > target:
		return binarySearch2(nums, rotatePos, len(nums)-1, target)
	default:
		return binarySearch2(nums, 0, rotatePos, target)
	}
}

// findRotatePosition 找出旋转的起始位置 未发生选择返回 0
func findRotatePosition(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return 0
	}

	// [8,9,3,4,5] => 9
	// [1,3,1,1,1] => 3
	// [1,1] => 0
	for lo, hi := 0, len(nums)-1; lo < hi; {
		mid := (lo + hi) / 2
		switch {
		case mid == len(nums)-1:
			return 0
		case nums[mid] > nums[mid+1]:
			return mid + 1
		case nums[lo] > nums[mid]:
			hi = mid - 1
		default:
			lo = mid + 1
		}
	}

	return 0
}

func binarySearch2(nums []int, lo, hi, target int) bool {
	for lo <= hi {
		mid := (lo + hi) / 2
		switch {
		case nums[mid] == target:
			return true
		case nums[mid] > target:
			hi = mid - 1
		case nums[mid] < target:
			lo = mid + 1
		}
	}

	return false
}
