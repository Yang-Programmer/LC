package medium

import (
	"leetcode/utils"
)

/**
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。



示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：true
示例 2：


输入：root = [1,2,2,3,3,null,null,4,4]
输出：false
示例 3：

输入：root = []
输出：true


提示：

树中的节点数在范围 [0, 5000] 内
-104 <= Node.val <= 104

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/balanced-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isBalanced(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	leftDepth := getMaxDepth(root.Left)
	rightDepth := getMaxDepth(root.Right)
	if absInt(leftDepth-rightDepth) < 1 {
		return true
	}
	return false
}

func getMaxDepth(node *utils.TreeNode) (depth int) {
	if node == nil {
		return
	}
	depth = max(getMaxDepth(node.Left), getMaxDepth(node.Right)) + 1
	return
}

func absInt(value int) int {
	if value < 0 {
		return 0 - value
	}
	return value
}
