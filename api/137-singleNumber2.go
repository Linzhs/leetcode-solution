package api

const intSize = 32 << (^uint(0) >> 63)

// 将所有数的各个位相加 再除以3取余 得到的位为只出现1次的数值的位
// 相当于分组后逐个消除
func singleNumber2(nums []int) (result int) {
	for i := 0; i < intSize; i++ {
		sum := 0
		for _, v := range nums {
			sum += (v >> i) & 1
		}
		result |= (sum % 3) << i
	}

	return
}
