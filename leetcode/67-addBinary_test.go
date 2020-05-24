package leetcode

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		n, k string
		want string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, v := range tests {
		got := addBinary(v.n, v.k)
		if got != v.want {
			t.Errorf("test case (%+v) want %s but got %s", v, v.want, got)
		}
	}
}
