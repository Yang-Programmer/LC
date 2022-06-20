package medium

import (
	"math"
)

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。

你可以认为每种硬币的数量是无限的。



示例1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0


提示：

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func coinChange(coins []int, amount int) int {
	l := len(coins)
	dp := make([]int, amount+1)
	for idx, _ := range dp {
		dp[idx] = math.MaxInt32
	}
	dp[0] = 0
	minInt := func(i, j int) int {
		if j < i {
			return j
		}
		return i
	}
	//dp[x]  零钱为x是最少硬币数量
	for i := 0; i < l; i++ {
		for j := coins[i]; j <= amount; j++ {
			// 右边 dp[j] 不选当前coin 继承上一轮dp值  dp[j-coins[i]]+1 选当前coin 上一轮dp值+1（次数）
			dp[j] = minInt(dp[j-coins[i]]+1, dp[j])
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
