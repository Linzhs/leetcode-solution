package leetcode

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {

	// 从右到左
	for m > 0 && n > 0 {
		switch {
		case nums1[m-1] >= nums2[n-1]:
			nums1[m+n-1] = nums1[m-1]
			m--
		case nums1[m-1] < nums2[n-1]:
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	if n > 0 {
		for ; n > 0; n-- {
			nums1[n-1] = nums2[n-1]
		}
	}
}
