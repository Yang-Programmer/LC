package easy

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。


示例 1：

输入：s = "()"
输出：true
示例2：

输入：s = "()[]{}"
输出：true
示例3：

输入：s = "(]"
输出：false
示例4：

输入：s = "([)]"
输出：false
示例5：

输入：s = "{[]}"
输出：true


提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isValid(s string) bool {
	result := false
	l := len(s)
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)
	if l == 0 {
		result = true
		return result
	}
	for i := 0; i < l; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && m[s[i]] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return result
		}
	}
	if len(stack) == 0 {
		result = true
	}
	return result
}
