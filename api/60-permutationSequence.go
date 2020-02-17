package api

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/permutation-sequence/discuss/22507/%22Explain-like-I'm-five%22-Java-Solution-in-O(n)
func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}

	factorial := make([]int, n+1)
	factorial[0] = 1

	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i // 阶乘 {1, 1, 2, 6, 24, ..., n!}
		nums[i-1] = i                     // 1~n 数值
	}

	// 从0开始算
	k--

	// generate kth permutation
	sb := strings.Builder{}
	for i := 1; i <= n; i++ {
		index := k / factorial[n-i]
		sb.WriteString(strconv.Itoa(nums[index]))
		nums = append(nums[:index], nums[index+1:]...) // 删除已经被选取作为排列的元素
		k -= index * factorial[n-i]
	}

	return sb.String()
}
