package leetcode

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) (result []string) {
	if len(nums) == 0 {
		return
	}

	for i, j := 0, 0; j < len(nums); j++ {
		if j < len(nums)-1 && nums[j]+1 == nums[j+1] {
			continue
		}

		switch i == j {
		case true:
			result = append(result, strconv.Itoa(nums[i]))
		default:
			result = append(result, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		}
		i = j + 1
	}

	return
}
