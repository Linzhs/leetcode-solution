package api

import (
	"testing"
)

func TestGetPermutation(t *testing.T) {
	tests := []struct {
		n, k int
		want string
	}{
		{4, 14, "3142"},
		{3, 3, "213"},
	}

	for _, v := range tests {
		got := getPermutation(v.n, v.k)
		if got != v.want {
			t.Errorf("test case (%+v) want %s but got %s", v, v.want, got)
		}
	}
}
