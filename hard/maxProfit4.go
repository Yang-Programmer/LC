package hard

/**
给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



示例 1：

输入：k = 2, prices = [2,4,1]
输出：2
解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2：

输入：k = 2, prices = [3,2,6,5,0,3]
输出：7
解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。


提示：

0 <= k <= 100
0 <= prices.length <= 1000
0 <= prices[i] <= 1000


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
类比 买卖股票3 发现 除了0以外，偶数就是卖出，奇数就是买入。
*/
func maxProfit4(k int, prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	dp := make([][]int, l)
	maxInt := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	for idx, _ := range dp {
		dp[idx] = make([]int, 2*k+1)
	}
	dp[0][0] = 0
	for i := 1; i < 2*k+1; i += 2 {
		dp[0][i] = 0 - prices[0]
	}

	for i := 1; i < l; i++ {
		for j := 0; j < 2*k-1; j += 2 {
			dp[i][j+1] = maxInt(dp[i-1][j]-prices[i], dp[i-1][j+1])
			dp[i][j+2] = maxInt(dp[i-1][j+1]+prices[i], dp[i-1][j+2])
		}
	}
	profit := 0
	for _, v := range dp[l-1] {
		if v >= profit {
			profit = v
		}
	}
	return profit
}
