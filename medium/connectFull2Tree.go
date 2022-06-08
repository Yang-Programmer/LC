package medium

import "container/list"

/**
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有next 指针都被设置为 NULL。



示例 1：



输入：root = [1,2,3,4,5,6,7]
输出：[1,#,2,3,#,4,5,6,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。
示例 2:

输入：root = []
输出：[]


提示：

树中节点的数量在[0, 212- 1]范围内
-1000 <= node.val <= 1000


进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/populating-next-right-pointers-in-each-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type ConnectNode struct {
	Val   int
	Left  *ConnectNode
	Right *ConnectNode
	Next  *ConnectNode
}

func connect(root *ConnectNode) *ConnectNode {
	if root == nil {
		return root
	}
	ret := make([][]*ConnectNode, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		curLen := queue.Len()
		tmp := make([]*ConnectNode, 0)
		for i := 0; i < curLen; i++ {
			top := queue.Remove(queue.Front()).(*ConnectNode)
			tmp = append(tmp, top)
			if top.Left != nil {
				queue.PushBack(top.Left)
			}
			if top.Right != nil {
				queue.PushBack(top.Right)
			}
		}
		ret = append(ret, tmp)
	}
	for idx, _ := range ret {
		for j := 0; j < len(ret[idx])-1; j++ {
			ret[idx][j].Next = ret[idx][j+1]
		}
	}
	return root
}
