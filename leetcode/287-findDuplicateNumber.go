package leetcode

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	p1, p2 := nums[0], nums[slow]
	for p1 != p2 {
		p1, p2 = nums[p1], nums[p2]
	}

	return p1
}
