package medium

import "strconv"

/**
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入'.' 来形成。你 不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。



示例 1：

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
示例 2：

输入：s = "0000"
输出：["0.0.0.0"]
示例 3：

输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]


提示：

1 <= s.length <= 20
s 仅由数字组成


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/restore-ip-addresses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func restoreIpAddresses(s string) []string {
	ret := make([]string, 0)
	l := len(s)
	if l == 0 {
		return ret
	}
	path := make([]string, 0)
	dote := 0
	var backtracking func(s string, index int)
	backtracking = func(s string, index int) {
		if dote > 3 {
			return
		}
		if dote == 3 {
			str := s[index:]
			if !isValid(str) {
				return
			}
			lp := len(path)
			ip := ""
			for j := 0; j < lp-1; j++ {
				ip = ip + path[j] + "."
			}
			ip += path[len(path)-1]
			ip += s[index:]
			ret = append(ret, ip)
			return
		}
		for i := index; i < l; i++ {
			str := s[index : i+1]
			if isValid(str) {
				path = append(path, s[index:i+1])
				dote++
			} else {
				break
			}
			backtracking(s, i+1)
			dote--
			path = path[:len(path)-1]
		}
	}
	backtracking(s, 0)
	return ret
}

func isValid(s string) bool {
	ret := true
	if len(s) == 0 {
		ret = false
		return ret
	}
	start := s[0] - '0'
	if start == 0 && len(s) != 1 {
		ret = false
		return ret
	}
	sInt, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		ret = false
		return ret
	}
	if sInt > 255 {
		ret = false
		return ret
	}
	return ret
}
