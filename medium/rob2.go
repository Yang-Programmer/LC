package medium

import "fmt"

/**
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。



示例1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
    偷窃到的最高金额 = 1 + 3 = 4 。
示例 3：

输入：nums = [1,2,3]
输出：3


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 1000
通过次数254,096提交次数580,359

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/house-robber-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
考虑两种情况
1、假设 第一个房间偷了，所以最后一个房间不能偷 那么可偷的范围就是 [0,l-2]
2、假设第一个房间没偷 那么最后一个房间就能偷  那么可偷的范围就是 [1,l-1]
取两者的最大值
*/
func rob2(nums []int) int {
	l := len(nums)

	maxInt := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return maxInt(nums[0], nums[1])
	}

	rob := func(start, end int) int {
		l := end - start + 1
		if l == 1 {
			return nums[start]
		}
		if l == 2 {
			return maxInt(nums[start], nums[end])
		}
		fmt.Println(l)
		dp := make([]int, l+1)
		dp[start] = nums[start]
		dp[start+1] = maxInt(nums[start], nums[start+1])
		for i := start + 2; i <= end; i++ {
			dp[i] = maxInt(dp[i-1], dp[i-2]+nums[i])
		}
		fmt.Println(dp)
		return dp[end]
	}
	v1, v2 := rob(0, l-2), rob(1, l-1)
	return maxInt(v1, v2)
}