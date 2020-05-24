package leetcode

func singleNumber3(nums []int) (result []int) {
	bitMark := 0
	for _, v := range nums {
		bitMark ^= v
	}

	diff := bitMark & (-bitMark)

	x := 0
	for _, v := range nums {
		if v&diff != 0 {
			x ^= v
		}
	}

	result = append(result, x, bitMark^x)

	return
}
