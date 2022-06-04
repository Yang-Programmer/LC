package easy

import (
	"fmt"
	"testing"
)

func TestImplementQueueUsingStack(t *testing.T) {
	myQueue := Constructor()
	myQueue.Push(1)          // queue is: [1]
	myQueue.Push(2)          // queue is: [1, 2] (leftmost is front of the queue)
	peek := myQueue.Peek()   // return 1
	pop := myQueue.Pop()     // return 1, queue is [2]
	empty := myQueue.Empty() // return false
	fmt.Println(peek, pop, empty)
}
