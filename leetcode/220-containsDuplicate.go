package leetcode

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}

	fnGetBucketId := func(x, bucketWidth int) int {
		if x < 0 {
			return (x+1)/bucketWidth - 1
		}
		return x / bucketWidth
	}

	fnAbs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var (
		bucketWidth  = t + 1             // 桶数 t==0时有一个桶
		rangeMapping = make(map[int]int) // 桶的区间 [0, t]、[t+1, 2t+1]...
	)
	for i, v := range nums {
		bucketId := fnGetBucketId(v, bucketWidth)

		if _, ok := rangeMapping[bucketId]; ok {
			// 桶中元素bucketElem和v大小小于t 且距离最多为k
			return true
		}

		// 跨桶比较
		if bucketElem, ok := rangeMapping[bucketId-1]; ok && fnAbs(bucketElem-v) < bucketWidth {
			return true
		}
		if bucketElem, ok := rangeMapping[bucketId+1]; ok && fnAbs(bucketElem-v) < bucketWidth {
			return true
		}

		rangeMapping[bucketId] = v
		// 一个桶只需存在一个元素
		if i >= k {
			delete(rangeMapping, fnGetBucketId(nums[i-k], bucketWidth))
		}
	}

	return false
}
