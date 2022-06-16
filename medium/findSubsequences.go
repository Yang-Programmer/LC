package medium

/**
给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。

数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。



示例 1：

输入：nums = [4,6,7,7]
输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
示例 2：

输入：nums = [4,4,3,2,1]
输出：[[4,4]]


提示：

1 <= nums.length <= 15
-100 <= nums[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/increasing-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findSubsequences(nums []int) [][]int {
	ret := make([][]int, 0)
	l := len(nums)
	if l == 0 {
		return ret
	}
	path := make([]int, 0)
	var backtracking func(nums []int, index int)
	backtracking = func(nums []int, index int) {
		if len(path) >= 2 {
			cp := make([]int, len(path))
			copy(cp, path)
			ret = append(ret, cp)
		}
		usedMap := make(map[int]struct{})
		for i := index; i < l; i++ {
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			if _, ok := usedMap[nums[i]]; ok {
				continue
			}
			path = append(path, nums[i])
			usedMap[nums[i]] = struct{}{}
			backtracking(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(nums, 0)
	return ret
}
