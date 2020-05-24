package leetcode

type wordTrieNode struct {
	val       byte
	isWord    bool
	charTable []*wordTrieNode
}

type WordDictionary struct {
	root *wordTrieNode
}

/** Initialize your data structure here. */
func ConstructorWordDictionary() (node WordDictionary) {
	node.root = &wordTrieNode{
		val:       ' ',
		isWord:    false,
		charTable: make([]*wordTrieNode, 26),
	}
	return
}

/** Adds a word into the data structure. */
func (p *WordDictionary) AddWord(word string) {
	node := p.root
	for _, v := range []byte(word) {
		if node.charTable[v-'a'] == nil {
			node.charTable[v-'a'] = &wordTrieNode{
				val:       v,
				isWord:    false,
				charTable: make([]*wordTrieNode, 26),
			}
		}
		node = node.charTable[v-'a']
	}

	node.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (p *WordDictionary) Search(word string) bool {
	return searchWord(p.root, []byte(word))
}

func searchWord(node *wordTrieNode, wordBytes []byte) bool {
	for i, v := range wordBytes {
		switch v == '.' {
		case true:
			for _, curNode := range node.charTable {
				if curNode != nil && searchWord(curNode, wordBytes[i+1:]) {
					return true
				}
			}
			return false
		default:
			curNode := node.charTable[v-'a']
			if curNode == nil || curNode.val != v {
				return false
			}
			node = curNode
		}
	}

	return node.isWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
