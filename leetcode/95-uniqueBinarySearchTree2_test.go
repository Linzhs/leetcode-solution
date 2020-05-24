package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	tests := []struct {
		n    int
		want []*TreeNode
	}{
		{n: 3},
	}

	for _, v := range tests {
		got := generateTrees(v.n)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("test case (%+v) want %+v but got %+v", v, v.want, got)
		}
	}
}
