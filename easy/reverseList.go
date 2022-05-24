package easy

import "leetcode/utils"

/**
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

https://leetcode.cn/problems/reverse-linked-list/
*/
func reverseList(head *utils.ListNode) *utils.ListNode {
	cur := head
	var pre *utils.ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

//递归版本
func reverseListV2(head *utils.ListNode) *utils.ListNode {
	return recursion(nil, head)
}
func recursion(pre, cur *utils.ListNode) *utils.ListNode {
	if cur == nil {
		return pre
	}
	tmp := cur.Next
	cur.Next = pre
	return recursion(cur, tmp)
}
