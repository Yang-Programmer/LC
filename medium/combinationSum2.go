package medium

import "sort"

/**
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。

candidates中的每个数字在每个组合中只能使用一次。

注意：解集不能包含重复的组合。



示例1:

输入: candidates =[10,1,2,7,6,1,5], target =8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例2:

输入: candidates =[2,5,2,1,2], target =5,
输出:
[
[1,2,2],
[5]
]


提示:

1 <=candidates.length <= 100
1 <=candidates[i] <= 50
1 <= target <= 30

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func combinationSum2(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	sum := 0
	var backtracking func(candidates []int, index int, target int)
	backtracking = func(candidates []int, index int, target int) {
		if sum > target {
			return
		}
		if sum == target {
			cp := make([]int, len(path))
			copy(cp, path)
			ret = append(ret, cp)
			return
		}
		l := len(candidates)
		for i := index; i < l; i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backtracking(candidates, i+1, target)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	sort.Ints(candidates)
	backtracking(candidates, 0, target)
	return ret
}
