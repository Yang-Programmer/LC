package easy

import (
	"container/list"
	"leetcode/utils"
)

/**
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。
*/
func minDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	depth := 0
	for queue.Len() > 0 {
		curLen := queue.Len()
		for i := 0; i < curLen; i++ {
			top := queue.Remove(queue.Front()).(*utils.TreeNode)
			if top.Left != nil {
				queue.PushBack(top.Left)
			}
			if top.Right != nil {
				queue.PushBack(top.Right)
			}
			if top.Left == nil && top.Right == nil {
				return depth
			}
		}
		depth++
	}
	return depth
}
