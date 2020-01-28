package api

import "strconv"

func isValidSudoku(board [][]byte) bool {
	var rows, cols, sudokus [9]int16

	for i, rowsBytes := range board {
		for j, vStr := range rowsBytes {
			if vStr == '.' {
				continue
			}

			v, _ := strconv.ParseUint(string(vStr), 10, 8)

			if (rows[i]>>(v-1))&0x01 == 1 {
				return false
			} else {
				rows[i] |= 1 << (v - 1)
			}

			if (cols[j]>>(v-1))&0x01 == 1 {
				return false
			} else {
				cols[j] |= 1 << (v - 1)
			}

			boxIndex := (i/3)*3 + j/3
			if (sudokus[boxIndex]>>(v-1))&0x01 == 1 {
				return false
			} else {
				sudokus[boxIndex] |= 1 << (v - 1)
			}
		}
	}

	return true
}
