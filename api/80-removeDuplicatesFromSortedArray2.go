package api

// sorted array
// 模拟有另一个数组的情况
// [0,0,1,1,1,1,2,3,3] ==>
// [0,0,1,1,2,3,3]
//
// [1,1,1,2,2,3] ==>
// [1,1,2,2,3]
func removeDuplicates2(nums []int) int {
	cur := 0
	for _, v := range nums {
		if cur < 2 {
			cur++
			continue
		}

		if v > nums[cur-2] {
			nums[cur] = v
			cur++
		}
	}

	return cur
}

func removeDuplicatesN(nums []int, n int) int {
	cur := 0
	for _, v := range nums {
		if cur < n {
			cur++
			continue
		}

		if v > nums[cur-n] {
			nums[cur] = v
			cur++
		}
	}

	return cur
}
