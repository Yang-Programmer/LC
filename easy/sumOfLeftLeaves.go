package easy

import (
	"container/list"
	"leetcode/utils"
)

/**
给定二叉树的根节点root，返回所有左叶子之和。



示例 1：



输入: root = [3,9,20,null,null,15,7]
输出: 24
解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
示例2:

输入: root = [1]
输出: 0


提示:

节点数在[1, 1000]范围内
-1000 <= Node.val <= 1000


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sum-of-left-leaves
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func sumOfLeftLeaves(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		curLen := queue.Len()
		for i := 0; i < curLen; i++ {
			top := queue.Remove(queue.Front()).(*utils.TreeNode)
			if top.Left != nil && top.Left.Left == nil && top.Left.Right == nil {
				ret += top.Left.Val
			}
			if top.Left != nil {
				queue.PushBack(top.Left)
			}
			if top.Right != nil {
				queue.PushBack(top.Right)
			}
		}
	}
	return ret
}

func sumOfLeftLeavesV2(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := sumOfLeftLeavesV2(root.Left)
	rightValue := sumOfLeftLeavesV2(root.Right)
	leavesValue := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leavesValue = root.Left.Val
	}
	return leftValue + rightValue + leavesValue
}
