package leetcode

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{name: "test case 1", args: args{nums: []int{1, 2, 3, 4}}, wantResult: []int{24, 12, 8, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := productExceptSelf(tt.args.nums); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("productExceptSelf() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
