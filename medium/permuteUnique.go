package medium

import "sort"

/**
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。



示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func permuteUnique(nums []int) [][]int {
	ret := make([][]int, 0)
	l := len(nums)
	if l == 0 {
		return ret
	}
	var backtracking func(nums []int, usedMap []bool)
	path := make([]int, 0)
	backtracking = func(nums []int, usedMap []bool) {
		if len(path) == l {
			cp := make([]int, l)
			copy(cp, path)
			ret = append(ret, cp)
			return
		}
		for i := 0; i < l; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				if !usedMap[i-1] {
					continue
				}
			}
			if !usedMap[i] {
				path = append(path, nums[i])
				usedMap[i] = true
				backtracking(nums, usedMap)
				path = path[:len(path)-1]
				usedMap[i] = false
			}
		}
	}
	sort.Ints(nums)
	backtracking(nums, make([]bool, l))
	return ret
}
