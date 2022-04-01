package medium

func LengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	right, result := -1, 0
	n := len(s)

	for i := 0; i < n; i++ {
		//每新向右移动一次窗口，重置窗口左边的元素的出现次数
		if i != 0 {
			delete(m, s[i-1])
		}
		for right+1 < n && m[s[right+1]] == 0 {
			//窗口右指针向右移，直到碰到元素出现次数大于1次
			m[s[right+1]] = m[s[right+1]] + 1
			right++
		}
		result = max(result, right-i+1)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
