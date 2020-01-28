package api

// 输入的整数范围为 1~3999
func intToRoman(num int) string {

	// 1000, 200, 300
	M := []string{"", "M", "MM", "MMM"}
	// 100, 200, 300...
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	// 10, 20, 30...
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	// 1, 2, 3...
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}
