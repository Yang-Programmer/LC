package medium

import "leetcode/utils"

/**
小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为root。

除了root之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

给定二叉树的root。返回在不触动警报的情况下，小偷能够盗取的最高金额。



示例 1:



输入: root = [3,2,3,null,3,null,1]
输出: 7
解释:小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
示例 2:



输入: root = [3,4,5,1,3,null,1]
输出: 9
解释:小偷一晚能够盗取的最高金额 4 + 5 = 9


提示：

树的节点数在[1, 104] 范围内
0 <= Node.val <= 104


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/house-robber-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func rob3(root *utils.TreeNode) int {
	maxInt := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	var robRange func(node *utils.TreeNode) [2]int
	robRange = func(node *utils.TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}
		left := robRange(node.Left)
		right := robRange(node.Right)
		//偷当前节点 则两个子节点都不能考虑
		robCurrent := node.Val + left[0] + right[0]
		//不偷当前节点 两个子节点都可以考虑（偷或者不偷 取最大值）
		notRobCurrent := maxInt(left[0], left[1]) + maxInt(right[0], right[1])
		return [2]int{notRobCurrent, robCurrent}
	}
	robMoney := robRange(root)
	return maxInt(robMoney[0], robMoney[1])
}
