package medium

import (
	"leetcode/utils"
)

/**
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。



示例 1：


输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]
示例 2：


输入：root = [1,2,3], targetSum = 5
输出：[]
示例 3：

输入：root = [1,2], targetSum = 0
输出：[]


提示：

树中节点总数在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/path-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func pathSum(root *utils.TreeNode, targetSum int) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	var traversal func(node *utils.TreeNode, count int, path []int)
	traversal = func(node *utils.TreeNode, count int, path []int) {
		if node.Left == nil && node.Right == nil && count == 0 {
			// 精髓 不拷贝的话这里存入了正确的值 后面修改path会影响这里的已存入的值
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			ret = append(ret, copyPath)
			return
		}
		if node.Left == nil && node.Right == nil && count != 0 {
			return
		}
		/*copyPath := make([]int, len(path))
		copy(copyPath, path)*/
		if node.Left != nil {
			count -= node.Left.Val
			path = append(path, node.Left.Val)
			//fmt.Println("left before: ", node.Left.Val, path, count, node.Left.Left, node.Left.Right)
			traversal(node.Left, count, path)
			count += node.Left.Val
			path = path[0 : len(path)-1]
			//fmt.Println("left after: ", node.Left.Val, path, count, node.Left.Left, node.Left.Right)
		}
		if node.Right != nil {
			count -= node.Right.Val
			path = append(path, node.Right.Val)
			//fmt.Println("right before: ", node.Right.Val, path, count, node.Right.Left, node.Right.Right)
			traversal(node.Right, count, path)
			count += node.Right.Val
			path = path[0 : len(path)-1]
			//fmt.Println("right after: ", node.Right.Val, path, count, node.Right.Left, node.Right.Right)
		}
	}
	traversal(root, targetSum-root.Val, []int{root.Val})
	return ret
}
