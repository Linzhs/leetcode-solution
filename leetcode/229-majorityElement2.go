package leetcode

// > n/3 至多只有2组
// 参考: https://leetcode-cn.com/problems/majority-element-ii/solution/169ti-sheng-ji-ban-xiang-jie-zhu-xing-jie-shi-tong/
func majorityElement2(nums []int) (result []int) {
	if len(nums) == 0 {
		return
	}

	var candidate1, candidate2, count1, count2 int
	for _, v := range nums {
		switch {
		case candidate1 == v:
			count1++
		case candidate2 == v:
			count2++
		case count1 == 0:
			candidate1 = v
			count1++
		case count2 == 0:
			candidate2 = v
			count2++
		default:
			count1--
			count2--
		}
	}

	count1 = 0
	count2 = 0
	for _, v := range nums {
		if v == candidate1 {
			count1++
			continue
		}
		if v == candidate2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		result = append(result, candidate1)
	}
	if count2 > len(nums)/3 {
		result = append(result, candidate2)
	}

	return
}
