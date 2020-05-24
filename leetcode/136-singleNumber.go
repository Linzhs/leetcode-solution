package leetcode

func singleNumber(nums []int) int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}

	return xor
}
