package easy

import (
	"leetcode/utils"
	"strconv"
)

/**
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。


示例 1：


输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]
示例 2：

输入：root = [1]
输出：["1"]


提示：

树中节点的数目在范围 [1, 100] 内
-100 <= Node.val <= 100

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func binaryTreePaths(root *utils.TreeNode) []string {
	ret := make([]string, 0)
	if root == nil {
		return ret
	}

	var traversal func(node *utils.TreeNode, s string)
	traversal = func(node *utils.TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			strVal := strconv.Itoa(node.Val)
			v := s + strVal
			ret = append(ret, v)
		}
		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			traversal(node.Left, s)
		}
		if node.Right != nil {
			traversal(node.Right, s)
		}
	}
	traversal(root, "")
	return ret
}
