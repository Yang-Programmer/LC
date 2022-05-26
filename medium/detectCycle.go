package medium

import (
	"leetcode/utils"
)

/**
给定一个链表的头节点 head，返回链表开始入环的第一个节点。如果链表无环，则返回null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func detectCycle(head *utils.ListNode) *utils.ListNode {
	result := new(utils.ListNode)
	cur := head
	pointerMap := make(map[*utils.ListNode]struct{})
	for cur != nil {
		pointerMap[cur] = struct{}{}
		if _, ok := pointerMap[cur.Next]; ok {
			result = cur.Next
			break
		}
		cur = cur.Next
	}
	return result
}

// 进阶：你是否可以使用 O(1) 空间解决此题？
// 快慢指针法 有点难
func detectCycleV2(head *utils.ListNode) *utils.ListNode {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			//无环
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
