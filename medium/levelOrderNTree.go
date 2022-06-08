package medium

import "container/list"

/**
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。



示例 1：



输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]
示例 2：



输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]


提示：

树的高度不会超过1000
树的节点总数在 [0,10^4] 之间

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/n-ary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Node struct {
	Val      int
	Children []*Node
}

func levelOrderNTree(root *Node) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		curLen := queue.Len()
		tmpArr := make([]int, 0)
		for i := 0; i < curLen; i++ {
			top := queue.Remove(queue.Front()).(*Node)
			tmpArr = append(tmpArr, top.Val)
			lenChildren := len(top.Children)
			for j := 0; j < lenChildren; j++ {
				queue.PushBack(top.Children[j])
			}
		}
		ret = append(ret, tmpArr)
	}
	return ret
}
