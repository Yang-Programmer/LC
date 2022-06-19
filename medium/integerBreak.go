package medium

/**
给定一个正整数n，将其拆分为 k 个 正整数 的和（k >= 2），并使这些整数的乘积最大化。

返回 你可以获得的最大乘积。



示例 1:

输入: n = 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例2:

输入: n = 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36。


提示:

2 <= n <= 58

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/integer-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	max := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
		}
	}
	return dp[len(dp)-1]
}
