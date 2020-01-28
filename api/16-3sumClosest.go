package api

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	s := sort.IntSlice(nums)
	sort.Sort(s)

	result := math.MaxInt64
	for i := 0; i < len(s)-2; i++ {
		for lo, hi := i+1, len(s)-1; lo < hi; {
			sum := s[i] + s[lo] + s[hi]
			if sum > target {
				hi--
			} else {
				lo++
			}
			if math.Abs(float64(sum-target)) < math.Abs(float64(result-target)) {
				result = sum
			}
		}
	}

	return result
}
