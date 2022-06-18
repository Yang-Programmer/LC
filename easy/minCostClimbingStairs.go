package easy

/**
给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。



示例 1：

输入：cost = [10,15,20]
输出：15
解释：你将从下标为 1 的台阶开始。
- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
总花费为 15 。
示例 2：

输入：cost = [1,100,1,1,1,100,1,1,100,1]
输出：6
解释：你将从下标为 0 的台阶开始。
- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
总花费为 6 。


提示：

2 <= cost.length <= 1000
0 <= cost[i] <= 999

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/min-cost-climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func minCostClimbingStairs(cost []int) int {
	minCost := 0
	l := len(cost)
	dp := make([]int, l)

	dp[0] = cost[0]
	dp[1] = cost[1]
	minInt := func(i, j int) int {
		if j < i {
			return j
		}
		return i
	}
	if l <= 2 {
		minCost = minInt(dp[0], dp[1])
		return minCost
	}
	for i := 3; i < l; i++ {
		dp[i] = minInt(dp[i-1], dp[i-2]) + cost[i]
	}
	return minInt(dp[l-1], dp[l-2])
}
