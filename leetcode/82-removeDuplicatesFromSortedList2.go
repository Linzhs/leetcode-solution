package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 虚拟一个头部 方便双指针遍历
	virtualHead := &ListNode{}
	virtualHead.Next = head

	// first指针指向重复节点的上一个节点
	for firstPointer, secondPointer := virtualHead, head; secondPointer != nil; {
		// 相同节点迭代前进
		for secondPointer.Next != nil && secondPointer.Val == secondPointer.Next.Val {
			secondPointer = secondPointer.Next
		}

		// first和second指针所指节点val不相等有两种情况
		// 1. 相差一到多个节点 则存在重复节点
		// 2. 相邻 则没有重复节点
		if firstPointer.Next != secondPointer { // 1
			firstPointer.Next = secondPointer.Next
			secondPointer = secondPointer.Next
			continue
		}

		// 2
		firstPointer = firstPointer.Next
		secondPointer = secondPointer.Next

	}

	return virtualHead.Next
}
