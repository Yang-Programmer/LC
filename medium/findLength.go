package medium

/**
给两个整数数组nums1和nums2，返回 两个数组中 公共的 、长度最长的子数组的长度。



示例 1：

输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
输出：3
解释：长度最长的公共子数组是 [3,2,1] 。
示例 2：

输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
输出：5


提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-length-of-repeated-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findLength(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	dp := make([][]int, l1+1)
	for idx, _ := range dp {
		dp[idx] = make([]int, l2+1)
	}
	for i := 0; i < l1+1; i++ {
		dp[i][0] = 0
	}
	for j := 0; j < l2+1; j++ {
		dp[j][0] = 0
	}
	result := 0
	// 可以看见如果以dp定义为i，j下标结束的相等子串长度的话 i和j得从1 开始 但是这样会导致nums1[0]和 nums2[0]无法统一处理
	//for i := 1; i < l1; i++ {
	//	for j := 1; j < l2; j++ {
	//		if nums1[i] == nums2[j] {
	//			dp[i][j] = dp[i-1][j-1] + 1
	//		}
	//		if dp[i][j] > result {
	//			result = dp[i][j]
	//		}
	//	}
	//}
	/*
		定义dp[i][j] 为下标[i-1] [j-1]的最长重复子串长度
	*/
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}

	return result
}
