package medium

import (
	"leetcode/utils"
)

/**
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗二叉树。



示例 1:


输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]
示例 2:

输入：inorder = [-1], postorder = [-1]
输出：[-1]


提示:

1 <= inorder.length <= 3000
postorder.length == inorder.length
-3000 <= inorder[i], postorder[i] <= 3000
inorder和postorder都由 不同 的值组成
postorder中每一个值都在inorder中
inorder保证是树的中序遍历
postorder保证是树的后序遍历

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func buildTree(inorder []int, postorder []int) *utils.TreeNode {
	var traversal func(inorder []int, postorder []int) *utils.TreeNode
	traversal = func(inorder []int, postorder []int) *utils.TreeNode {
		if len(postorder) == 0 || len(inorder) == 0 {
			return nil
		}
		if len(postorder) == 1 {
			return &utils.TreeNode{
				Val: postorder[0],
			}
		}
		// 后序遍历最后一个元素一定是当前递归的根节点
		rootValue := postorder[len(postorder)-1]
		root := &utils.TreeNode{
			Val: rootValue,
		}
		//找到中序遍历中该节点的index
		inorderIdx := -1
		for idx, v := range inorder {
			if v == rootValue {
				inorderIdx = idx
				break
			}
		}
		// 根据中序遍历的index 拆分中序遍历找到左右节点
		leftInorder := inorder[:inorderIdx]
		rightInorder := inorder[inorderIdx+1:]
		leftPostorder := postorder[:inorderIdx]
		rightPostorder := postorder[inorderIdx : len(postorder)-1]
		root.Left = traversal(leftInorder, leftPostorder)
		root.Right = traversal(rightInorder, rightPostorder)
		return root
	}
	node := traversal(inorder, postorder)
	return node
}
func buildTreePreorderInorder(preorder []int, inorder []int) *utils.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var traversal func(preorder []int, inorder []int) *utils.TreeNode
	traversal = func(preorder []int, inorder []int) *utils.TreeNode {
		if len(preorder) == 1 {
			return &utils.TreeNode{Val: preorder[0]}
		}
		preorderValue := preorder[0]
		root := &utils.TreeNode{Val: preorderValue}
		inorderIdx := -1
		for idx, v := range inorder {
			if v == preorderValue {
				inorderIdx = idx
				break
			}
		}
		leftInorder := inorder[:inorderIdx]
		rightInorder := inorder[inorderIdx+1:]

		leftPreorder := preorder[1 : 1+len(leftInorder)]
		rightPreorder := preorder[1+len(leftPreorder):]
		root.Left = traversal(leftPreorder, leftInorder)
		root.Right = traversal(rightPreorder, rightInorder)
		return root
	}
	return traversal(preorder, inorder)
}
