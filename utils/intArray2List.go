package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Array2ListInt(list []int) *ListNode {
	dummyHead := new(ListNode)
	cur := dummyHead
	for _, item := range list {
		tmp := &ListNode{
			Val:  item,
			Next: nil,
		}
		cur.Next = tmp
		cur = cur.Next
	}
	return dummyHead.Next
}

func PrintLinkListElements(list *ListNode) {
	cur := list
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
