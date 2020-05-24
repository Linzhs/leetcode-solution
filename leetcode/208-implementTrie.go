package leetcode

type trieNode struct {
	val rune
	isWord bool
	childTable []*trieNode
}

type Trie struct {
	root *trieNode
}


/** Initialize your data structure here. */
func Constructor1() (node Trie) {
	root := &trieNode {
		val: ' ',
		childTable: make([]*trieNode, 26),
	}

	node.root = root

	return
}


/** Inserts a word into the trie. */
func (p *Trie) Insert(word string)  {
	node := p.root
	for _, v := range word {
		if node.childTable[v-'a'] == nil {
			node.childTable[v-'a'] = &trieNode {
				val: v,
				childTable: make([]*trieNode, 26),
			}
		}
		node = node.childTable[v-'a']
	}

	node.isWord = true
}


/** Returns if the word is in the trie. */
func (p *Trie) Search(word string) bool {
	node := p.root
	for _, v := range word {
		curNode := node.childTable[v-'a']
		if curNode == nil || curNode.val != v {
			return false
		}
		node = curNode
	}

	return node.isWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (p *Trie) StartsWith(prefix string) bool {
	node := p.root
	for _, v := range prefix {
		curNode := node.childTable[v-'a']
		if curNode == nil || curNode.val != v {
			return false
		}
		node = curNode
	}

	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */