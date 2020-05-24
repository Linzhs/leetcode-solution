package leetcode

func maxProfit1(prices []int) (result int) {
	if len(prices) <= 1 {
		return
	}

	dp := prices[0]
	for _, v := range prices {
		sub := v - dp
		if sub > result {
			result = sub
		}

		if dp > v {
			dp = v
		}
	}

	return
}
