package api

import (
	"reflect"
	"testing"
)

func TestMergeInters(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		{[][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	}

	for _, v := range tests {
		got := merge(v.intervals)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("test case (%v) want %v but got %v", v, v.want, got)
		}
	}
}
