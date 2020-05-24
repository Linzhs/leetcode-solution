package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		prevIdx, ok := m[v]
		if ok && i-prevIdx <= k {
			return true
		}
		m[v] = i
	}

	return false
}
