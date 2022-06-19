package medium

/**
给你一个 只包含正整数 的 非空 数组nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。



示例 1：

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
示例 2：

输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。


提示：

1 <= nums.length <= 200
1 <= nums[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/partition-equal-subset-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func canPartition(nums []int) bool {
	sum := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	maxInt := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	target := sum / 2
	/**
	01背包
	背包的体积为sum / 2
	背包要放入的商品（集合里的元素）重量为 元素的数值，价值也为元素的数值
	背包如果正好装满，说明找到了总和为 sum / 2 的子集。
	背包中每一个元素是不可重复放入。
	本题中 dp[j]表示容量为j的背包最大价值
	*/
	dp := make([]int, target+1)
	for i := 0; i < l; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = maxInt(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}
