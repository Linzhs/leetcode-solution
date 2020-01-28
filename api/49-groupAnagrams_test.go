package api

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, want: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
	}

	for _, v := range tests {
		got := groupAnagrams(v.strs)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("test case(%v) want %v but got %v", v, v.want, got)
		}
	}
}
