package pkg

import (
	"fmt"
	"leetcode-solution/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPaths(t *testing.T) {
	grid := [][]bool{
		{false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, true, false},
		{false, false, false, false, true, false, false, false},
		{true, false, true, false, false, true, false, false},
		{false, false, true, false, false, false, false, false},
		{false, false, false, true, true, false, true, false},
		{false, true, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false},
	}

	fmt.Println(countPathsV2(grid))
}

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 3},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(api.minimumTotalV1(triangle))
	fmt.Println(api.minimumTotalV2(triangle))
}

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	got := api.maxProduct(nums)
	assert.Equal(t, 6, got)
}

func TestCoinsChange(t *testing.T) {
	tests := []struct {
		arr    []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, api.coinChange(test.arr, test.amount))
	}
}

func TestLargestRectangleArea(t *testing.T) {
	s := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(api.largestRectangleArea(s))
}

func TestLIS(t *testing.T) {
	fmt.Println(api.lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"babad", "aba"},
		{"cbbd", "bb"},
		{"", ""},
		{"a", "a"},
		{"aa", "aa"},
		{"ab", "b"},
		{"aba", "aba"},
		{"aba", "aba"},
		{"sbb", "bb"},
		{"sfb", "b"},
		{"aaabaaaa", "aaabaaa"},
	}

	for _, test := range tests {
		got := api.longestPalindrome(test.s)
		if got != test.want {
			t.Fatalf("input '%s' want '%s' but got '%s'", test.s, test.want, got)
		}
	}
}
