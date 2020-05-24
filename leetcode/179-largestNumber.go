package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i]+p[j] > p[j]+p[i] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p StringSlice) Sort() { sort.Sort(p) }

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return "0"
	}

	var sSlice StringSlice
	for _, v := range nums {
		sSlice = append(sSlice, strconv.Itoa(v))
	}

	if len(sSlice) == 0 {
		return "0"
	}
	sSlice.Sort()

	if sSlice[0] == "0" {
		return "0"
	}

	return strings.Join(sSlice, "")
}
