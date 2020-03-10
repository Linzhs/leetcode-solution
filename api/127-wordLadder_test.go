package api

import (
	"testing"
)

func TestWordLadder(t *testing.T) {
	if ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}) != 5 {
		t.Errorf("fail")
	}
}
