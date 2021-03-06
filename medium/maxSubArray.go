package medium

/**
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。



示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23


提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxSubArray(nums []int) int {
	ret := nums[0]
	count := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		count += nums[i]
		if count > ret {
			ret = count
		}
		if count < 0 {
			//说明加了个负数 当前值是个负数不要了
			count = 0
		}
	}
	return ret
}
func maxSubArrayDp(nums []int) int {
	l := len(nums)
	maxInt := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	//dp[i]表示下标为i时连续子数组最大的值
	//我们可以考虑{nums}[i] 单独成为一段还是加入 f(i-1)对应的那一段，
	//这取决于 nums[i] 和 f(i−1)+nums[i] 的大小，我们希望获得一个比较大的，

	dp := make([]int, l)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < l; i++ {
		dp[i] = maxInt(dp[i-1]+nums[i], nums[i])
		result = maxInt(result, dp[i])
	}
	return result
}
