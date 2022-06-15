package medium

/**
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。





示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]


提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func letterCombinations(digits string) []string {
	ret := make([]string, 0)
	if len(digits) == 0 {
		return ret
	}
	path := make([]byte, 0)
	digitsMap := [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var backtracking func(digits string, index int)
	backtracking = func(digits string, index int) {
		if index == len(digits) {
			cp := make([]byte, len(digits))
			copy(cp, path)
			ret = append(ret, string(path))
			return
		}
		//取出输入字符串的索引位置的字符串并转化为数字
		num := digits[index] - '0'
		//根据数字取出键盘上对应的字符串
		num2Str := digitsMap[num]
		lNum2Str := len(num2Str)
		for i := 0; i < lNum2Str; i++ {
			//遍历键盘字符串
			path = append(path, num2Str[i])
			backtracking(digits, index+1) //递归处理下一个输入字符串代表的数字
			path = path[:len(path)-1]     //回溯
		}
	}
	backtracking(digits, 0)
	return ret
}
