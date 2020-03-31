package api

func twoSum2(numbers []int, target int) []int {
	for lo, hi := 0, len(numbers)-1; lo < hi; {
		v := numbers[lo] + numbers[hi]
		switch {
		case v > target:
			hi--
		case v < target:
			lo++
		default:
			return []int{lo, hi}
		}
	}

	return nil
}
