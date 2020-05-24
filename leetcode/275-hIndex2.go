package leetcode

func hIndex2(citations []int) int {
	i := 0
	for i < len(citations) && citations[len(citations) - i - 1] > i {
		i++
	}

	return i
}
