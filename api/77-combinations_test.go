package api

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want [][]int
	}{
		{4, 2, [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}},
	}

	for _, test := range tests {
		got := combine(test.n, test.k)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("test case (n:%d, k:%d) want %v but got %v", test.n, test.k, test.want, got)
		}
	}
}
