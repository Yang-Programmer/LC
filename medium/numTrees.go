package medium

/**
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。



示例 1：


输入：n = 3
输出：5
示例 2：

输入：n = 1
输出：1


提示：

1 <= n <= 19


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/unique-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
题目要求是计算不同二叉搜索树的个数。为此，我们可以定义两个函数：

G(n): 长度为 n 的序列能构成的不同二叉搜索树的个数。

F(i, n): 以 i 为根、序列长度为 n 的不同二叉搜索树个数 (1 <= i <= n)(1≤i≤n)。



*/
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	//dp[i] 等于1到i的节点遍历中 以每个节点x为根节点，[1,x) 长度x-1 为左节点 (x+1,i]为右节点长度i-x 的dp值之和
	/*
		F(i,n)=G(i−1)⋅G(n−i)
	*/
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[len(dp)-1]
}
