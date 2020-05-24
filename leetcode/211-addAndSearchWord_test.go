package leetcode

import "testing"

// ["WordDictionary","addWord","addWord","addWord","addWord","search","search","addWord","search","search","search","search","search","search"]
// [[],["at"],["and"],["an"],["add"],["a"],[".at"],["bat"],[".at"],["an."],["a.d."],["b."],["a.d"],["."]]
func TestWordDictionary_Search(t *testing.T) {
	dict := ConstructorWordDictionary()
	dict.AddWord("at")
	dict.AddWord("and")
	dict.AddWord("an")
	dict.AddWord("add")
	dict.Search("a")
	dict.Search(".at")
	dict.AddWord("bat")
	dict.Search(".at")
	dict.Search("an.")
	dict.Search("a.d.")
	dict.Search("b.")
	dict.Search("a.d")
	dict.Search(".")
}
