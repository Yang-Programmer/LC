package medium

import (
	"container/list"
	"leetcode/utils"
)

/**
给定一个二叉树的 根节点 root，请找出该二叉树的最底层最左边节点的值。

假设二叉树中至少有一个节点。



示例 1:



输入: root = [2,1,3]
输出: 1
示例 2:



输入: [1,2,3,4,null,5,6,null,null,7]
输出: 7


提示:

二叉树的节点个数的范围是 [1,104]
-231<= Node.val <= 231- 1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/find-bottom-left-tree-value
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findBottomLeftValue(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	arr := make([][]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		curLen := queue.Len()
		tmp := make([]int, 0)
		for i := 0; i < curLen; i++ {
			top := queue.Remove(queue.Front()).(*utils.TreeNode)
			tmp = append(tmp, top.Val)
			if top.Left != nil {
				queue.PushBack(top.Left)
			}
			if top.Right != nil {
				queue.PushBack(top.Right)
			}
		}
		arr = append(arr, tmp)
	}
	ret := arr[len(arr)-1]
	return ret[0]
}
