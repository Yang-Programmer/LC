package medium

import (
	"container/list"
	"leetcode/utils"
)

/**
给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）



示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[15,7],[9,20],[3]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 将从上至下的层序遍历数组反转一下就可以了
func levelOrderBottom(root *utils.TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		curLen := queue.Len()
		tmpArr := make([]int, 0)
		for i := 0; i < curLen; i++ {
			top := queue.Remove(queue.Front()).(*utils.TreeNode)
			tmpArr = append(tmpArr, top.Val)
			if top.Left != nil {
				queue.PushBack(top.Left)
			}
			if top.Right != nil {
				queue.PushBack(top.Right)
			}
		}
		ret = append(ret, tmpArr)
	}
	reverserArr := func(arr [][]int) {
		i, j := 0, len(arr)-1
		for i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	reverserArr(ret)
	return ret
}
