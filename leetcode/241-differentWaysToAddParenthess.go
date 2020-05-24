package leetcode

import "strconv"

func diffWaysToCompute(input string) (result []int) {
	if len(input) == 0 {
		return
	}

	for i, v := range input {
		if v != '+' && v != '-' && v != '*' {
			continue
		}

		leftPartResult := diffWaysToCompute(input[:i])
		rightPartResult := diffWaysToCompute(input[i+1:])

		for _, lpr := range leftPartResult {
			for _, rpr := range rightPartResult {
				switch v {
				case '+':
					result = append(result, lpr+rpr)
				case '-':
					result = append(result, lpr-rpr)
				case '*':
					result = append(result, lpr*rpr)
				}
			}
		}
	}

	if len(result) == 0 {
		v, err := strconv.Atoi(input)
		if err == nil {
			result = append(result, v)
		}
	}

	return
}
