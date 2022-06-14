package easy

import "leetcode/utils"

/**
给定二叉搜索树（BST）的根节点root和一个整数值val。

你需要在 BST 中找到节点值等于val的节点。 返回以该节点为根的子树。 如果节点不存在，则返回null。



示例 1:



输入：root = [4,2,7,1,3], val = 2
输出：[2,1,3]
Example 2:


输入：root = [4,2,7,1,3], val = 5
输出：[]


提示：

数中节点数在[1, 5000]范围内
1 <= Node.val <= 107
root是二叉搜索树
1 <= val <= 107


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/search-in-a-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
二叉搜索树是一个有序树：

若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
它的左、右子树也分别为二叉搜索树

*/
func searchBST(root *utils.TreeNode, val int) *utils.TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return nil
}
