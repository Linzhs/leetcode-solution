package leetcode

// 可以使用前n项和求和公式 sum=n(n+1)/2
func missingNumber(nums []int) int {
	missing := len(nums)
	for i, v := range nums {
		missing ^= i ^ v
	}
	return missing
}
