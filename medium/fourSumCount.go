package medium

/**
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0


示例 1：

输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释：
两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
示例 2：

输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
输出：1




来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/4sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum1 := make(map[int]int)
	result := 0
	for _, item1 := range nums1 {
		for _, item2 := range nums2 {
			sum1[item1+item2]++
		}
	}
	for _, item1 := range nums3 {
		for _, item2 := range nums4 {
			result += sum1[-item1-item2]
		}
	}
	return result
}
