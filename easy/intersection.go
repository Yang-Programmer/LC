package easy

/**

349. 两个数组的交集
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的


提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/
func intersection(nums1 []int, nums2 []int) []int {
	history := make(map[int]struct{})
	result := make([]int, 0)
	for _, item := range nums1 {
		history[item] = struct{}{}
	}
	for _, item := range nums2 {
		if _, ok := history[item]; ok {
			result = append(result, item)
			delete(history, item)
		}
	}
	return result
}
