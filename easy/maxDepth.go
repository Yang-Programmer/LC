package easy

import (
	"container/list"
	"leetcode/utils"
)

/**
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明:叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度3 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	depth := 0
	for queue.Len() > 0 {
		curLen := queue.Len()
		depth++
		for i := 0; i < curLen; i++ {
			top := queue.Remove(queue.Front()).(*utils.TreeNode)
			if top.Left != nil {
				queue.PushBack(top.Left)
			}
			if top.Right != nil {
				queue.PushBack(top.Right)
			}
		}
	}
	return depth
}
