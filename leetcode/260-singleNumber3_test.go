package leetcode

import (
	"reflect"
	"testing"
)

func Test_singleNumber3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{name: "test case 1", args: args{nums: []int{6, 7, 8, 8}}, wantResult: []int{6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := singleNumber3(tt.args.nums); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("singleNumber3() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
