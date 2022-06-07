package easy

import "leetcode/utils"

/**
给你二叉树的根节点 root ，返回它节点值的前序遍历。



示例 1：


输入：root = [1,null,2,3]
输出：[1,2,3]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：


输入：root = [1,2]
输出：[1,2]
示例 5：


输入：root = [1,null,2]
输出：[1,2]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100


进阶：递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 迭代法 使用栈实现二叉树的前序遍历
func preorderTraversal(root *utils.TreeNode) []int {
	stack := make([]*utils.TreeNode, 0)
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack = append(stack, root)
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		ret = append(ret, top.Val)
		stack = stack[:len(stack)-1]
		right := top.Right
		left := top.Left
		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
	}
	return ret
}
