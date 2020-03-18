package api

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		v, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, v)
			continue
		}

		n1 := stack[len(stack)-2]
		n2 := stack[len(stack)-1]
		stack = stack[:len(stack)-2]

		switch token {
		case "+":
			stack = append(stack, n1+n2)
		case "-":
			stack = append(stack, n1-n2)
		case "*":
			stack = append(stack, n1*n2)
		case "/":
			stack = append(stack, n1/n2)
		}
	}

	return stack[0]
}
