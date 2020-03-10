package api

type randomPointerNode struct {
	val    int
	next   *randomPointerNode
	random *randomPointerNode
}

func copyRandomList(node *randomPointerNode) *randomPointerNode {
	m := make(map[*randomPointerNode]*randomPointerNode)

	return copyRandomListWithMap(m, node)
}

func copyRandomListWithMap(m map[*randomPointerNode]*randomPointerNode, node *randomPointerNode) *randomPointerNode {
	if node == nil {
		return nil
	}

	if curNode, ok := m[node]; ok {
		return curNode
	}

	newNode := &randomPointerNode{val: node.val}
	m[node] = newNode

	newNode.next = copyRandomListWithMap(m, node.next)
	newNode.random = copyRandomListWithMap(m, node.random)

	return newNode
}

func copyRandomListV2(head *randomPointerNode) *randomPointerNode {
	// 依据next指针遍历节点并拷贝每个节点到自身的下一个节点
	// A->B->C->D ====> A->A'->B->B'->C->C'->D->D'
	for curNode := head; curNode != nil; {
		newNode := &randomPointerNode{val: curNode.val}

		newNode.next = curNode.next
		curNode.next = newNode
		curNode = newNode.next
	}

	// 为拷贝的节点设置正确的random指针
	for curNode := head; curNode != nil; curNode = curNode.next.next {
		if curNode.random != nil {
			curNode.next.random = curNode.random.next
		}
	}

	// 分离链表
	newListHead := head.next
	for oldNode, newNode := head, head.next; oldNode != nil; {
		oldNode.next = oldNode.next.next

		if newNode.next != nil {
			newNode.next = newNode.next.next
		}

		oldNode = oldNode.next
		newNode = newNode.next
	}

	return newListHead
}
