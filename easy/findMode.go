package easy

import "leetcode/utils"

/**
给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。

如果树中有不止一个众数，可以按 任意顺序 返回。

假定 BST 满足如下定义：

结点左子树中所含节点的值 小于等于 当前节点的值
结点右子树中所含节点的值 大于等于 当前节点的值
左子树和右子树都是二叉搜索树


示例 1：


输入：root = [1,null,2,2]
输出：[2]
示例 2：

输入：root = [0]
输出：[0]


提示：

树中节点的数目在范围 [1, 104] 内
-105 <= Node.val <= 105

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/find-mode-in-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findMode(root *utils.TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	base := 0
	count := 0
	maxCount := 0
	var traversal func(root *utils.TreeNode)
	traversal = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		if root.Val == base {
			count++
		} else {
			base, count = root.Val, 1
		}
		if count > maxCount {
			maxCount = count
			ret = []int{base}
		} else if count == maxCount {
			ret = append(ret, base)
		}
		traversal(root.Right)
	}
	traversal(root)
	return ret
}
