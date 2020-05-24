package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{name: "3-7", args: args{k: 3, n: 7}, wantResult: [][]int{{1, 2, 4}}},
		{name: "3-9", args: args{k: 3, n: 9}, wantResult: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("combinationSum3() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
