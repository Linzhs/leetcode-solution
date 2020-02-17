package api

// group
func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// p0 跟踪0的最右边界
	// p2 跟踪2的最左边界
	// cur跟踪当前元素
	for p0, cur, p2 := 0, 0, len(nums)-1; cur <= p2; {
		switch nums[cur] {
		case 0:
			nums[cur], nums[p0] = nums[p0], nums[cur]
			p0++
			cur++
		case 1:
			cur++
		case 2:
			nums[cur], nums[p2] = nums[p2], nums[cur]
			p2--
		}
	}
}
