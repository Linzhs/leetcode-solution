package leetcode

var uglySlice []int

func init() {
	uglySlice = make([]int, 1690)
	uglySlice[0] = 1

	i2, i3, i5 := 0, 0, 0

	for i := 1; i < 1690; i++ {
		ugly := min(min(2*uglySlice[i2], 3*uglySlice[i3]), 5*uglySlice[i5])
		uglySlice[i] = ugly

		if ugly == uglySlice[i2]*2 {
			i2++
		}
		if ugly == uglySlice[i3]*3 {
			i3++
		}
		if ugly == uglySlice[i5]*5 {
			i5++
		}
	}
}

func nthUglyNumber(n int) int {
	return uglySlice[n-1]
}
