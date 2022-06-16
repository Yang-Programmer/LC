package medium

/**
你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。



示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]


提示：

1 <= s.length <= 16
s 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func partition(s string) [][]string {
	ret := make([][]string, 0)
	path := make([]string, 0)
	l := len(s)
	if l == 0 {
		return ret
	}
	var backtracking func(s string, index int)
	backtracking = func(s string, index int) {
		if index >= l {
			cp := make([]string, len(path))
			copy(cp, path)
			ret = append(ret, cp)
			return
		}
		for i := index; i < l; i++ {
			if isPalindStr(s[index : i+1]) {
				path = append(path, s[index:i+1])
			} else {
				continue
			}
			backtracking(s, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(s, 0)
	return ret
}

func isPalindStr(s string) (ret bool) {
	l := len(s)
	i, j := 0, l-1
	ret = true
	for i < j {
		if s[i] != s[j] {
			ret = false
			break
		}
		i++
		j--
	}
	return ret
}
