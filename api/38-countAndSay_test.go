package api

import "testing"

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{4, "1211"},
	}

	for _, v := range tests {
		got := countAndSay(v.n)
		if got != v.want {
			t.Errorf("test case (%+v) got %q but want %q", v, got, v.want)
		}
	}
}
