package easy

/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."


限制：

0 <= s 的长度 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/ti-huan-kong-ge-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func replaceSpace(s string) string {
	//统计空格的数量
	sb := []byte(s)
	oldLen := len(s)
	var count int
	for _, v := range sb {
		if v == ' ' {
			count++
		}
	}
	empty := make([]byte, count*2)
	sb = append(sb, empty...)
	newLen := len(sb)
	left, right := oldLen-1, newLen-1
	for left >= 0 {
		if sb[left] == ' ' {
			sb[right] = '0'
			sb[right-1] = '2'
			sb[right-2] = '%'
			left--
			right = right - 3
		} else {
			sb[right] = sb[left]
			left--
			right--
		}
	}
	return string(sb)
}
