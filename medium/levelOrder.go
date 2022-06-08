package medium

import (
	"container/list"
	"leetcode/utils"
)

/**
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：
输入：root = [1]
输出：[[1]]
示例 3：
输入：root = []
输出：[]

提示：

树中节点数目在范围 [0, 2000] 内
-1000 <= Node.val <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func levelOrder(root *utils.TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		tmpArr := make([]int, 0)
		curLen := queue.Len()
		for i := 0; i < curLen; i++ {
			top := queue.Remove(queue.Front()).(*utils.TreeNode)
			if top.Left != nil {
				queue.PushBack(top.Left)
			}
			if top.Right != nil {
				queue.PushBack(top.Right)
			}
			tmpArr = append(tmpArr, top.Val)
		}
		ret = append(ret, tmpArr)
	}
	return ret
}
