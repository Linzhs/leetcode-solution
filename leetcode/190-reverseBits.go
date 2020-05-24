package leetcode

func reverseBits(num uint32) (result uint32) {
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= num & 1
		num >>= 1
	}

	return
}
