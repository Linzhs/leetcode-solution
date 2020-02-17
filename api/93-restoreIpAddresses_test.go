package api

import (
	"reflect"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
	}

	for _, v := range tests {
		got := restoreIpAddresses(v.s)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("test case (%+v) want %+v but got %+v", v, v.want, got)
		}
	}
}
