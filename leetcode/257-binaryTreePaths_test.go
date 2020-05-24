package leetcode

import (
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{name: "test case 1", args: args{root: nil},},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := binaryTreePaths(tt.args.root); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("binaryTreePaths() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}