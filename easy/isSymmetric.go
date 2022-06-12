package easy

import (
	"leetcode/utils"
)

/**

给你一个二叉树的根节点 root ， 检查它是否轴对称。



示例 1：


输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：


输入：root = [1,2,2,null,3,null,3]
输出：false


提示：

树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/symmetric-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isSymmetric(root *utils.TreeNode) bool {
	var compare func(left, right *utils.TreeNode) bool
	compare = func(left, right *utils.TreeNode) bool {
		if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if left == nil && right == nil {
			return true
		} else if left.Val != right.Val {
			return false
		}
		outSide := compare(left.Left, right.Right)
		inSide := compare(left.Right, right.Left)
		return outSide && inSide
	}
	if root == nil {
		return false
	}
	return compare(root.Left, root.Right)
}

//迭代法
func isSymmetricV2(root *utils.TreeNode) bool {
	if root == nil {
		return false
	}
	queue := make([]*utils.TreeNode, 0)
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
