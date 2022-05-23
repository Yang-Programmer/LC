package medium

/**
螺旋矩阵

给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。



示例 1：


输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：

输入：n = 1
输出：[[1]]


提示：

1 <= n <= 20

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/spiral-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func generateMatrix(n int) [][]int {
	//golang 初始化二维数组
	result := make([][]int, n)
	for idx, _ := range result {
		result[idx] = make([]int, n)
	}
	num, target := 1, n*n // 填充数
	left, right, top, button := 0, n-1, 0, n-1
	for num <= target {
		// step 1 左上=>右上
		for column := left; column <= right; column++ {
			result[top][column] = num
			num++
		}
		top++
		// step 2 右上 => 右下
		for row := top; row <= button; row++ {
			result[row][right] = num
			num++
		}
		right--
		//step 3 右下 => 左下
		for column := right; column >= left; column-- {
			result[button][column] = num
			num++
		}
		button--
		//step 4 右下 => 左上
		for row := button; row >= top; row-- {
			result[row][left] = num
			num++
		}
		left++
	}
	return result
}
