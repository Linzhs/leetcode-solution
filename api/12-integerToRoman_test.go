package api

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{n: 3, want: "III"},
		{n: 4, want: "IV"},
		{n: 9, want: "IX"},
		{n: 58, want: "LVIII"},
		{n: 1994, want: "MCMXCIV"},
	}

	for _, v := range tests {
		got := intToRoman(v.n)
		if got != v.want {
			t.Errorf("test case (%#+v) got %s but want %s", v, got, v.want)
		}
	}
}
