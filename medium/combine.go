package medium

/**
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。



示例 1：

输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
示例 2：

输入：n = 1, k = 1
输出：[[1]]


提示：

1 <= n <= 20
1 <= k <= n

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	queue := make([]int, 0)
	var backtracking func(n int, idx int, k int)
	backtracking = func(n int, idx int, k int) {
		if len(queue) == k {
			cp := make([]int, k)
			copy(cp, queue)
			ret = append(ret, cp)
			return
		}
		for i := idx; i <= n; i++ {
			queue = append(queue, i)
			backtracking(n, i+1, k)
			queue = queue[:len(queue)-1]
		}
	}
	backtracking(n, 1, k)
	return ret
}

func combineV2(n int, k int) [][]int {
	ret := make([][]int, 0)
	queue := make([]int, 0)
	var backtracking func(n int, idx int, k int)
	backtracking = func(n int, idx int, k int) {
		if len(queue) == k {
			cp := make([]int, k)
			copy(cp, queue)
			ret = append(ret, cp)
			return
		}
		// 还需要(k-size)个元素 循环内再把i加进去 最终值不能超过n i+(k-size)-1<n => i<= n-(k-size)+1
		for i := idx; i <= n-(k-len(queue))+1; i++ {
			queue = append(queue, i)
			backtracking(n, i+1, k)
			queue = queue[:len(queue)-1]
		}
	}
	backtracking(n, 1, k)
	return ret
}
