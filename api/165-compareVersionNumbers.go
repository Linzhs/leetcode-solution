package api

import (
	"strconv"
	"strings"
)

func compareVersion(version1, version2 string) int {
	var (
		vSlice1 = strings.Split(version1, ".")
		vSlice2 = strings.Split(version2, ".")
	)

	maxLen := len(vSlice1)
	if len(vSlice1) < len(vSlice2) {
		maxLen = len(vSlice2)
	}

	for i := 0; i < maxLen; i++ {
		v1 := 0
		if i < len(vSlice1) {
			v1, _ = strconv.Atoi(vSlice1[i])
		}

		v2 := 0
		if i < len(vSlice2) {
			v2, _ = strconv.Atoi(vSlice2[i])
		}

		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
	}

	return 0
}
