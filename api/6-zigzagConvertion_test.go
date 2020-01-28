package api

import "testing"

func TestZiaZagConvert(t *testing.T) {
	tests := []struct {
		s    string
		n    int
		want string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	}

	for _, test := range tests {
		got := convert(test.s, test.n)
		if got != test.want {
			t.Errorf("test case (s:%s n:%d) want %s but got %s",
				test.s, test.n, test.want, test.s)
		}
	}
}
