package leetcode

import (
	"testing"
)

func TestQueueConstructor(t *testing.T) {
	q := QueueConstructor()
	q.Push(1)
	q.Push(2)
	q.Pop()
	q.Push(3)
	q.Push(4)
	q.Pop()
	q.Peek()
}
