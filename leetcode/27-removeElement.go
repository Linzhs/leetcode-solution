package leetcode

func removeElement(nums []int, val int) int {
	for lo, hi := 0, len(nums)-1; lo < hi; {

		for lo < hi && nums[lo] != val {
			lo++
		}
		for hi > lo && nums[hi] == val {
			hi--
		}

		if lo < hi && nums[lo] == val {
			nums[lo], nums[hi] = nums[hi], nums[lo]
		}
	}

	for i, v := range nums {
		if v == val {
			return i
		}
		if i == len(nums)-1 && v != val {
			return len(nums)
		}
	}

	return 0
}
