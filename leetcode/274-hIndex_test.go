package leetcode

import "testing"

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{name: "test case 1", args: args{citations: []int{3, 0, 6, 1, 5}}, wantResult: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := hIndex(tt.args.citations); gotResult != tt.wantResult {
				t.Errorf("hIndex() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
