package leetcode

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{name: "test1", args: args{nums: []int{0, 1, 2, 4, 5, 7}}, wantResult: []string{"0->2", "4->5", "7"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := summaryRanges(tt.args.nums); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("summaryRanges() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
