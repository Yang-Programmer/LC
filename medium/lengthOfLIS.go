package medium

/**
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1


提示：

1 <= nums.length <= 2500
-104 <= nums[i] <= 104


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func lengthOfLIS(nums []int) int {
	l := len(nums)
	//dp[i] 表示下标为i(必须包含i)的元素最长的递增子串
	/**
	遍历[0:i)元素j 取最大的长度
	*/
	maxInt := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	dp := make([]int, l)
	for idx, _ := range dp {
		dp[idx] = 1
	}
	result := 1
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1) //dp[j]+1 加1是指加上nums[i]的这个元素
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result

}
