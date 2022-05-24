package medium

import (
	"fmt"
	"testing"
)

func TestDesignLinkedList(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(38)
	linkedList.AddAtTail(66)
	linkedList.AddAtTail(61)
	linkedList.AddAtTail(76)
	linkedList.AddAtTail(26)
	linkedList.AddAtTail(37)
	linkedList.AddAtTail(8)
	linkedList.DeleteAtIndex(5)
	linkedList.AddAtHead(4)
	linkedList.AddAtHead(45)
	linkedList.Get(4)
	linkedList.AddAtTail(85)
	linkedList.AddAtTail(93)
	linkedList.AddAtIndex(10, 23)
	linkedList.AddAtTail(21)
	linkedList.AddAtHead(52)
	linkedList.AddAtHead(15)
	linkedList.AddAtHead(47)
	num := linkedList.Get(12)
	fmt.Println(num, "-------")
	addr := linkedList.Next
	for addr != nil {
		fmt.Println(addr.Val)
		addr = addr.Next
	}
}
