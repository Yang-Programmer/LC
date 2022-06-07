package easy

import "leetcode/utils"

/**
给定一个二叉树的根节点 root ，返回 它的 中序遍历 。



示例 1：


输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func inorderTraversal(root *utils.TreeNode) []int {
	ret := make([]int, 0)
	var traversal func(node *utils.TreeNode)
	traversal = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		ret = append(ret, node.Val)
		traversal(node.Right)
	}

	traversal(root)
	return ret
}
