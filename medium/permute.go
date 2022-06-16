package medium

/**
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]


提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	l := len(nums)
	if l == 0 {
		return ret
	}
	path := make([]int, 0)
	var backtracking func(nums []int, usedMap map[int]struct{})

	backtracking = func(nums []int, usedMap map[int]struct{}) {
		if len(path) == l {
			cp := make([]int, l)
			copy(cp, path)
			ret = append(ret, cp)
			return
		}
		for i := 0; i < l; i++ {
			if _, ok := usedMap[nums[i]]; ok {
				continue
			}
			path = append(path, nums[i])
			usedMap[nums[i]] = struct{}{}
			backtracking(nums, usedMap)
			path = path[:len(path)-1]
			delete(usedMap, nums[i])
		}
	}
	backtracking(nums, make(map[int]struct{}))
	return ret
}
