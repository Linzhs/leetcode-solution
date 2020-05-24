package leetcode

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		stack []int
		sign  = '+'
		num   int
	)
	for i, v := range s {
		if isDigital(v) {
			num = num*10 + int(v-'0')
		}

		if (!isDigital(v) && v != ' ') || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top*num)
			case '/':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top/num)
			}

			sign = v
			num = 0
		}
	}

	var result int
	for _, v := range stack {
		result += v
	}

	return result
}

func isDigital(r rune) bool {
	return r >= '0' && r <= '9'
}
