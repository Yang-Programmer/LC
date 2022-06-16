package medium

import "sort"

/**
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。



示例 1：

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/subsets-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func subsetsWithDup(nums []int) [][]int {
	ret := make([][]int, 0)
	l := len(nums)
	if l == 0 {
		return ret
	}
	path := make([]int, 0)
	var backtracking func(nums []int, index int)

	backtracking = func(nums []int, index int) {
		cp := make([]int, len(path))
		copy(cp, path)
		ret = append(ret, cp)
		if index >= l {
			return
		}
		for i := index; i < l; i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtracking(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	sort.Ints(nums)
	backtracking(nums, 0)
	return ret
}
