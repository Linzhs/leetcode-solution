package api

import "sort"

type twoDiIntSlice [][]int

func (p twoDiIntSlice) Len() int      { return len(p) }
func (p twoDiIntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p twoDiIntSlice) Less(i, j int) bool {
	return p[i][0] < p[j][0] || (p[i][0] == p[j][0] && p[i][1] < p[j][1])
}

// [[1,3],[2,6],[8,10],[15,18]]
// [[1,4],[4,5]]
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Sort(twoDiIntSlice(intervals))

	var result [][]int
	for _, v := range intervals {
		if len(result) == 0 {
			result = append(result, v)
			continue
		}

		lastSlice := result[len(result)-1]
		if lastSlice[1] < v[0] {
			result = append(result, v)
			continue
		}

		lastSlice[1] = max(lastSlice[1], v[1])
	}

	return result
}
