package leetcode

import "strings"

func convert(s string, numRows int) string {

	if len(s) <= 1 || numRows > len(s) || numRows == 1 {
		return s
	}

	var (
		strArr = make([]string, min(len(s), numRows), min(len(s), numRows))
		downOp bool
		row    int
	)
	for _, c := range s {
		strArr[row] += string(c)
		if row == 0 || row == numRows-1 {
			downOp = !downOp
		}
		if downOp {
			row++
		} else {
			row--
		}
	}

	return strings.Join(strArr, "")
}
