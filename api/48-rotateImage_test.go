package api

import (
	"reflect"
	"testing"
)

func TestRotatedImage(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			}},
	}

	for _, v := range tests {
		rotate(v.matrix)
		if !reflect.DeepEqual(v.want, v.matrix) {
			t.Errorf("test case (%+v) got %v but want %v", v, v.matrix, v.want)
		}
	}
}
