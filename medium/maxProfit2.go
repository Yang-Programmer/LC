package medium

/**
给你一个整数数组 prices ，其中prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候最多只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润。



示例 1：

输入：prices = [7,1,5,3,6,4]
输出：7
解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
    随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
     总利润为 4 + 3 = 7 。
示例 2：

输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
    总利润为 4 。
示例3：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0 。


提示：

1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
动态规划转移方程
dp[i][0] 表示第i天手里没有股票时的最大现金
dp[i][1] 表示第i天手里有股票时的最大现金
对于dp[i][0] 1、第i-1天手里也没有股票 dp[i-1][0] 2、第i-1天有股票选择在第i天卖出 dp[i-1][1]+prices[i]
对于dp[i][1] 1、第i-1天手里也有股票 dp[i-1][1] 2、第i-1没有股票 在第i天买入的 dp[i-1][0]-prices[i]

动态规划初始化
dp[0][0] 第1天不持有股票 手里现金等于0
dp[0][1] 第1天持有股票   手里现金等于-prices[0]
*/
func maxProfit2(prices []int) int {
	l := len(prices)
	dp := make([][]int, l)
	for idx, _ := range dp {
		dp[idx] = make([]int, 2)
	}
	maxInt := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	dp[0][0] = 0
	dp[0][1] = 0 - prices[0]
	for i := 1; i < l; i++ {
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = maxInt(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[l-1][0]
}
