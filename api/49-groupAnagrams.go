package api

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	if len(strs) == 1 {
		return [][]string{strs}
	}

	m := make(map[string][]string)

	for _, v := range strs {
		s := strings.Split(v, "")
		sort.Sort(sort.StringSlice(s))
		sortedString := strings.Join(s, "")
		m[sortedString] = append(m[sortedString], v)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}
