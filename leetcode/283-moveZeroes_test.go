package leetcode

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test case 1", args: args{nums: []int{1, 0}}, want: []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.want, tt.args.nums) {
				t.Errorf("got %+v", tt.args.nums)
			}
		})
	}
}
