package medium

import (
	"container/list"
	"leetcode/utils"
)

/**
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。



示例 1:



输入:[1,2,3,null,5,null,4]
输出:[1,3,4]
示例 2:

输入:[1,null,3]
输出:[1,3]
示例 3:

输入:[]
输出:[]


提示:

二叉树的节点个数的范围是 [0,100]
-100<= Node.val <= 100


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-right-side-view
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//层序遍历的时候，判断是否遍历到单层的最后面的元素，如果是，就放进result数组中，随后返回result就可以了。
func rightSideView(root *utils.TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		curLen := queue.Len()
		for i := 0; i < curLen; i++ {
			top := queue.Remove(queue.Front()).(*utils.TreeNode)
			if i == curLen-1 {
				ret = append(ret, top.Val)
			}
			if top.Left != nil {
				queue.PushBack(top.Left)
			}
			if top.Right != nil {
				queue.PushBack(top.Right)
			}
		}
	}
	return ret
}
