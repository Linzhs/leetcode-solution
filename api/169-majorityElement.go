package api

func majorityElement(nums []int) int {
	var result, cnt int
	for _, v := range nums {
		switch {
		case cnt == 0:
			result = v
			cnt++
		case result == v:
			cnt++
		case result != v:
			cnt--
		}
	}

	return result
}
