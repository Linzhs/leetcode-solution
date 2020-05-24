package leetcode

import (
	"testing"
)

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint32
	}{
		{name:"testCase 1", args:args{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := reverseBits(tt.args.num); gotResult != tt.wantResult {
				t.Errorf("reverseBits() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}