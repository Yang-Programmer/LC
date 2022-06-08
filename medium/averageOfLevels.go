package medium

import (
	"container/list"
	"leetcode/utils"
)

/**
给定一个非空二叉树的根节点root, 以数组的形式返回每一层节点的平均值。与实际答案相差10-5 以内的答案可以被接受。



示例 1：



输入：root = [3,9,20,null,null,15,7]
输出：[3.00000,14.50000,11.00000]
解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
因此返回 [3, 14.5, 11] 。
示例 2:



输入：root = [3,9,20,15,7]
输出：[3.00000,14.50000,11.00000]


提示：

树中节点数量在[1, 104] 范围内
-231<= Node.val <= 231- 1


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/average-of-levels-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func averageOfLevels(root *utils.TreeNode) []float64 {
	ret := make([]float64, 0)
	if root == nil {
		return ret
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		curLen := queue.Len()
		sum := 0
		for i := 0; i < curLen; i++ {
			top := queue.Remove(queue.Front()).(*utils.TreeNode)
			sum += top.Val
			if top.Left != nil {
				queue.PushBack(top.Left)
			}
			if top.Right != nil {
				queue.PushBack(top.Right)
			}
		}
		avg := float64(sum) / float64(curLen)
		ret = append(ret, avg)
	}
	return ret
}
