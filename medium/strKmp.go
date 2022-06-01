package medium

/**
实现strStr()函数。

给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。

说明：

当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当needle是空字符串时我们应当返回 0 。这与 C 语言的strstr()以及 Java 的indexOf()定义相符。



示例 1：

输入：haystack = "hello", needle = "ll"
输出：2
示例 2：

输入：haystack = "aaaaa", needle = "bba"
输出：-1


提示：

1 <= haystack.length, needle.length <= 104
haystack 和 needle 仅由小写英文字符组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

// KMP 算法 O(len(haystack) + len(needle))
/**
参考文章：
https://www.cnblogs.com/tangzhengyue/p/4315393.html
https://www.cnblogs.com/yjiyjige/p/3263858.html
*/
func strStr(haystack string, needle string) int {
	targetIdx := -1
	lN := len(needle)
	lH := len(haystack)
	next := getNext(needle)
	i, j := 0, 0
	for i < lH && j < lN {
		if j == -1 || haystack[i] == needle[j] {
			j++
			i++
		} else {
			j = next[j]
		}
	}
	if j == lN {
		targetIdx = i - j
	}
	return targetIdx
}

func getNext(needle string) []int {
	nLen := len(needle)
	next := make([]int, nLen)
	next[0] = -1
	j, k := 0, -1
	for j < nLen-1 {
		if k == -1 || needle[k] == needle[j] {
			if needle[j+1] == needle[k+1] {
				next[j+1] = next[k+1]
			} else {
				next[j+1] = k + 1
			}
			j++
			k++
		} else {
			k = next[k]
		}
	}
	return next
}
