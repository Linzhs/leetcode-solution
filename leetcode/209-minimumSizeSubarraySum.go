package leetcode

import "math"

// 这里使用双指针的方式来实现
// 使用left记录当前连续子数组的起始位置
// 当连续子数组的sum >= s时, 则可以计算出连续子数组的
// 长度是否为最小, 并移动left
func minSubArrayLen(s int, nums []int) int {
	left, sum, result := 0, 0, math.MaxInt64
	for i, v := range nums {
		sum += v
		for sum >= s {
			result = min(result, i-left+1)
			// 移动窗口 需要将最左边的元素移除窗口
			sum -= nums[left]
			left++
		}
	}

	if result == math.MaxInt64 {
		return 0
	}

	return result
}
