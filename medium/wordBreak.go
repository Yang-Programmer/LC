package medium

/**
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。



示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
    注意，你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false


提示：

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s 和 wordDict[i] 仅有小写英文字母组成
wordDict 中的所有字符串 互不相同

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/word-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	//dp[x] 长度为x的字符串是否为满足条件的字符串
	dp := make([]bool, l+1)
	dp[0] = true
	for i := 1; i <= l; i++ {
		//长度为i的下标为i-1
		for j := i - 1; j >= 0; j-- {
			//长度为i的字符串为满足条件则存在某一个下标j [0,j-1]（dp[j]）满足 且[j,i-1]满足
			//dp[j]长度为j的字符串是否为满足条件的字符串下标为[0,j-1]
			subStr := s[j:i]
			if wordMap[subStr] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[l]
}
