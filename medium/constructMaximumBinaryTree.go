package medium

import "leetcode/utils"

/**
给定一个不重复的整数数组nums 。最大二叉树可以用下面的算法从nums 递归地构建:

创建一个根节点，其值为nums 中的最大值。
递归地在最大值左边的子数组前缀上构建左子树。
递归地在最大值 右边 的子数组后缀上构建右子树。
返回nums 构建的 最大二叉树 。



示例 1：


输入：nums = [3,2,1,6,0,5]
输出：[6,3,5,null,2,0,null,null,1]
解释：递归调用如下所示：
- [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
    - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
        - 空数组，无子节点。
        - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
            - 空数组，无子节点。
            - 只有一个元素，所以子节点是一个值为 1 的节点。
    - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
        - 只有一个元素，所以子节点是一个值为 0 的节点。
        - 空数组，无子节点。
示例 2：


输入：nums = [3,2,1]
输出：[3,null,2,null,1]


提示：

1 <= nums.length <= 1000
0 <= nums[i] <= 1000
nums 中的所有整数 互不相同

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func constructMaximumBinaryTree(nums []int) *utils.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var traversal func(nums []int) *utils.TreeNode
	traversal = func(nums []int) *utils.TreeNode {
		// 终止条件
		if len(nums) == 0 {
			return nil
		}
		if len(nums) == 1 {
			return &utils.TreeNode{Val: nums[0]}
		}
		//找出最大值和其下标
		max := -1
		maxIdx := -1
		for idx, v := range nums {
			if v > max {
				max = v
				maxIdx = idx
			}
		}
		left := nums[:maxIdx]
		right := nums[maxIdx+1:]
		root := &utils.TreeNode{
			Val:   max,
			Left:  traversal(left),
			Right: traversal(right),
		}
		return root
	}
	return traversal(nums)
}
