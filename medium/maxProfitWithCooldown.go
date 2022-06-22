package medium

/**
给定一个整数数组prices，其中第prices[i]表示第i天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



示例 1:

输入: prices = [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
示例 2:

输入: prices = [1]
输出: 0


提示：

1 <= prices.length <= 5000
0 <= prices[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxProfitWithCooldown(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	//dp[i][0~2]表示第i天 结束（注意是第i天结束）后处于不同状态下的剩余金钱
	/**
		0、 第i天结束后持有股票
		1、 第i天结束后未持有股票且处于冷冻期
		2、 第i天结束后未持有股票且不处于冷冻期
		dp[i][0] 1)第i天买入股票 则第i-1天一定是未持有股票且不处于冷冻期的 dp[i-1][2]-prices[i]
	             2)当天什么都没做 但是i-1天持有股票 dp[i-1][0]
	    dp[i][1] 一定是第i天卖掉股票注意dp[i]的含义是第i天结束（如果是卖掉股票的第二天的话那么结束后就不是冷冻期了）
	             dp[i-1][0]+prices[i]+prices[i]
	    dp[i][2] 那么i-1一定是未持有股票的（如果i-1买入了由于冷冻期 第i天不能把卖掉）
	             dp[i-1][1]（i-1天卖掉i-1天结束后处于冷冻期）和dp[i-1][2]都可能存在
	*/
	dp := make([][]int, l)
	maxInt := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	for idx, _ := range dp {
		dp[idx] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 1; i < l; i++ {
		dp[i][0] = maxInt(dp[i-1][2]-prices[i], dp[i-1][0])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = maxInt(dp[i-1][1], dp[i-1][2])
	}
	return maxInt(dp[l-1][1], dp[l-1][2])
}
