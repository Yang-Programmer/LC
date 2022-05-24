package medium

/**
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val和next。val是当前节点的值，next是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性prev以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第index个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为val的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第index个节点之前添加值为val 的节点。如果index等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引index 有效，则删除链表中的第index 个节点。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/design-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	dummyHead := MyLinkedList{}
	return dummyHead
}

func (this *MyLinkedList) Get(index int) int {
	var target int = -1
	cur := this.Next
	linkIndex := 0
	for cur != nil {
		if index == linkIndex {
			target = cur.Val
			break
		}
		linkIndex++
		cur = cur.Next
	}
	return target
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &MyLinkedList{
		Val: val,
	}
	newNode.Next = this.Next
	this.Next = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	idx := this
	for idx != nil && idx.Next != nil {
		idx = idx.Next
	}
	newNode := &MyLinkedList{
		Val: val,
	}
	idx.Next = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.Next
	linkIdx := 1 //虚拟头节点的好处
	for cur != nil {
		if index == linkIdx {
			nNode := new(MyLinkedList)
			nNode.Val = val
			nNode.Next = cur.Next
			cur.Next = nNode
			break
		} else {
			cur = cur.Next
			linkIdx++
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.Next
	linkIdx := 1
	if index == 0 {
		if cur != nil {
			this.Next = cur.Next
		}
		return
	}
	for cur != nil && cur.Next != nil {
		if index == linkIdx {
			cur.Next = cur.Next.Next
			break
		} else {
			cur = cur.Next
			linkIdx++
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
