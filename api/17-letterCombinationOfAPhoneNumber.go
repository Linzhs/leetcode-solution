package api

var phone = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var result []string
	backtrack(&result, "", digits)
	return result
}

func backtrack(result *[]string, combination, nextDigits string) {
	if len(nextDigits) == 0 {
		*result = append(*result, combination)
		return
	}

	digit := string(nextDigits[0])
	letters := phone[digit]
	for _, v := range letters {
		backtrack(result, combination+string(v), string(nextDigits[1:]))
	}
}
