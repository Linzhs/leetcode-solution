package api

import "strings"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	s := make([]int, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := int((num1[i] - '0') * (num2[j] - '0'))
			sum := mul + s[i+j+1]
			s[i+j] += sum / 10
			s[i+j+1] = sum % 10
		}
	}

	b := strings.Builder{}
	for _, v := range s {
		if b.Len() == 0 && v == 0 {
			continue
		}
		b.WriteRune(rune(v) + '0')
	}

	return b.String()
}
