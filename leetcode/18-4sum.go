package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) == 0 || len(nums) < 4 {
		return nil
	}

	var result [][]int

	s := sort.IntSlice(nums)
	s.Sort()

	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			continue
		}
		for j := i + 1; j < len(s)-1; j++ {
			if j > i+1 && s[j] == s[j-1] {
				continue
			}
			lo, hi := j+1, len(s)-1
			for lo < hi {
				sum := s[i] + s[j] + s[lo] + s[hi]
				if sum > target {
					hi--
					continue
				}
				if sum < target {
					lo++
					continue
				}

				result = append(result, []int{s[i], s[j], s[lo], s[hi]})
				for lo < hi && s[lo] == s[lo+1] {
					lo++
				}
				for hi > lo && s[hi] == s[hi-1] {
					hi--
				}
				lo++
				hi--
			}
		}
	}

	return result
}
