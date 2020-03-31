package api

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	nodeStack []*TreeNode
}

func ConstructorBstIterator(root *TreeNode) BSTIterator {
	bst := BSTIterator{}

	bst.nextValidNode(root)

	return bst
}

func (p *BSTIterator) nextValidNode(root *TreeNode) {
	for cur := root; cur != nil; cur = cur.Left {
		p.nodeStack = append(p.nodeStack, cur)
	}
}

/** @return the next smallest number */
func (p *BSTIterator) Next() int {
	elem := p.nodeStack[len(p.nodeStack)-1]
	p.nodeStack = p.nodeStack[:len(p.nodeStack)-1]

	p.nextValidNode(elem.Right)

	return elem.Val
}

/** @return whether we have a next smallest number */
func (p *BSTIterator) HasNext() bool {
	return len(p.nodeStack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
