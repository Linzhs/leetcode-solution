package leetcode

type MyQueue struct {
	Stack1 []int
	Stack2 []int
}

/** Initialize your data structure here. */
func QueueConstructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.Stack1 = append(q.Stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	q.Peek()

	top := q.Stack2[len(q.Stack2)-1]
	q.Stack2 = q.Stack2[:len(q.Stack2)-1]
	return top
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if len(q.Stack2) == 0 {
		for len(q.Stack1) > 0 {
			top := q.Stack1[len(q.Stack1)-1]
			q.Stack1 = q.Stack1[:len(q.Stack1)-1]
			q.Stack2 = append(q.Stack2, top)
		}
	}
	return q.Stack2[len(q.Stack2)-1]
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.Stack1) == 0 && len(q.Stack2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
