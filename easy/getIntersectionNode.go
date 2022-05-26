package easy

import "leetcode/utils"

/**
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/
func getIntersectionNode(headA, headB *utils.ListNode) *utils.ListNode {
	//双指针法
	if headA == nil || headB == nil {
		return nil
	}
	curA := headA
	curB := headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}
		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}
