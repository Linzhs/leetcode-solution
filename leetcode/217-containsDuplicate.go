package leetcode

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Sort(sort.IntSlice(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}
