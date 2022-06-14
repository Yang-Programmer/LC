package easy

import (
	"leetcode/utils"
	"math"
)

/**
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

差值是一个正数，其数值等于两值之差的绝对值。



示例 1：


输入：root = [4,2,6,1,3]
输出：1
示例 2：


输入：root = [1,0,48,null,null,12,49]
输出：1


提示：

树中节点的数目范围是 [2, 104]
0 <= Node.val <= 105


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-absolute-difference-in-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func getMinimumDifference(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	arr := make([]int, 0)
	var traversal func(root *utils.TreeNode)
	traversal = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		arr = append(arr, root.Val)
		traversal(root.Right)
	}
	traversal(root)
	l := len(arr)
	if l < 2 {
		return 0
	}
	ret := math.MaxInt
	for i := 1; i < l; i++ {
		ret = min(ret, arr[i]-arr[i-1])
	}
	return ret
}
func min(i, j int) int {
	if j < i {
		return j
	}
	return i
}
