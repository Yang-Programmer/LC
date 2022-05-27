package easy

/**
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。



示例 1：

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
示例 2：

输入：n = 2
输出：false


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/happy-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//有点傻逼 偏门
func isHappy(n int) bool {
	//如果没有数学证明 你怎么知道 无重复出现的sum值即代表其最终sum一定等于1呢？
	result := true
	m := make(map[int]struct{})
	for n != 1 {
		sum := getSum(n)
		if _, ok := m[sum]; ok {
			result = false
			break
		}
		m[sum] = struct{}{}
		n = sum
	}
	return result
}

func getSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
