package medium

/**
找出所有相加之和为n 的k个数的组合，且满足下列条件：

只使用数字1到9
每个数字最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。



示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
解释:
1 + 2 + 4 = 7
没有其他符合的组合了。
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
解释:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
没有其他符合的组合了。
示例 3:

输入: k = 4, n = 1
输出: []
解释: 不存在有效的组合。
在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。


提示:

2 <= k <= 9
1 <= n <= 60

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/combination-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func combinationSum3(k int, n int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	sum := 0
	var backtracking func(k int, n int, sum int, index int)
	backtracking = func(k int, n int, sum int, index int) {
		if sum > n {
			return
		}
		if len(path) == k {
			if sum == n {
				cp := make([]int, k)
				copy(cp, path)
				ret = append(ret, cp)
				return
			}
		}
		for i := index; i <= 9; i++ {
			path = append(path, i)
			sum += i
			backtracking(k, n, sum, i+1)
			path = path[:len(path)-1]
			sum -= i
		}
	}
	backtracking(k, n, sum, 1)
	return ret
}
