package api

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	var (
		strSlice []string
		carry    uint8
	)
	for aIter, bIter := len(a)-1, len(b)-1; aIter >= 0 || bIter >= 0; {
		sum := carry
		if aIter >= 0 {
			sum += a[aIter] - '0'
			aIter--
		}
		if bIter >= 0 {
			sum += b[bIter] - '0'
			bIter--
		}

		strSlice = append(strSlice, strconv.Itoa(int(sum%2)))
		carry = sum / 2
	}

	if carry > 0 {
		strSlice = append(strSlice, strconv.Itoa(int(carry)))
	}

	// reverse
	for lo, hi := 0, len(strSlice)-1; lo < hi; {
		strSlice[lo], strSlice[hi] = strSlice[hi], strSlice[lo]
		lo++
		hi--
	}

	return strings.Join(strSlice, "")
}
