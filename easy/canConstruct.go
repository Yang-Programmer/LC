package easy

/**
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。



示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false
示例 2：

输入：ransomNote = "aa", magazine = "ab"
输出：false
示例 3：

输入：ransomNote = "aa", magazine = "aab"
输出：true


提示：

1 <= ransomNote.length, magazine.length <= 105
ransomNote 和 magazine 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/ransom-note
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	magazineCharMap := [26]int{}
	for _, item := range magazine {
		magazineCharMap[item-'a']++
	}
	for _, item := range ransomNote {
		magazineCharMap[item-'a']--
		if magazineCharMap[item-'a'] < 0 {
			return false
		}
	}
	return true
}