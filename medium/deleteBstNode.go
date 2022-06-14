package medium

import "leetcode/utils"

/**
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的key对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。


示例 1:



输入：root = [5,3,6,2,4,null,7], key = 3
输出：[5,4,6,2,null,null,7]
解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
另一个正确答案是 [5,2,6,null,4,null,7]。


示例 2:

输入: root = [5,3,6,2,4,null,7], key = 0
输出: [5,3,6,2,4,null,7]
解释: 二叉树不包含值为 0 的节点
示例 3:

输入: root = [], key = 0
输出: []


提示:

节点数的范围[0, 104].
-105<= Node.val <= 105
节点值唯一
root是合法的二叉搜索树
-105<= key <= 105


进阶： 要求算法时间复杂度为O(h)，h 为树的高度。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/delete-node-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func deleteBSTNode(root *utils.TreeNode, key int) *utils.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			// 左右子节点为空（根节点） 删除根节点右节点
			return nil
		} else if root.Left == nil && root.Right != nil {
			// 左节点为空 右节点不为空 右节点补位
			return root.Right
		} else if root.Left != nil && root.Right == nil {
			// 左节点不为空 右节点为空 左节点补位
			return root.Left
		} else {
			// 左右节点都不为空  将左节点放在右节点的最左边的左子节点上 右节点成为新的root
			right := root.Right
			cur := right
			for cur.Left != nil {
				//找到右节点的最左左节点
				cur = cur.Left
			}
			cur.Left = root.Left //将当前节点的左节点放在右节点的最左
			return right
		}
	}
	if root.Val > key {
		root.Left = deleteBSTNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteBSTNode(root.Right, key)
	}
	return root
}
