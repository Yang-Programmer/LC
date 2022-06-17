package hard

import "strings"

/**
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的n皇后问题 的解决方案。

每一种解法包含一个不同的n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。



示例 1：


输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：[["Q"]]


提示：

1 <= n <= 9

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/n-queens
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func solveNQueens(n int) [][]string {
	ret := make([][]string, 0)
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	if n == 0 {
		return ret
	}
	var backtracking func(n int, row int)
	backtracking = func(n int, row int) {
		if row == n {
			tmp := make([]string, n)
			for idx, strArr := range chessboard {
				tmp[idx] = strings.Join(strArr, "")
			}
			ret = append(ret, tmp)
			return
		}
		for col := 0; col < n; col++ {
			if isValidQueens(row, col, n, chessboard) {
				chessboard[row][col] = "Q"
				backtracking(n, row+1)
				chessboard[row][col] = "."
			}
		}
	}
	backtracking(n, 0)
	return ret
}
func isValidQueens(row, col, n int, chessboard [][]string) bool {
	//检查当前位置这一列上面有没有出现过Queen
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	//检查 / 斜线是否出现过Queue 45度角
	i, j := row-1, col+1
	for i >= 0 && j < n {
		if chessboard[i][j] == "Q" {
			return false
		}
		i--
		j++
	}
	// 检查 \ 斜线是否出现过Queen 135度角
	p, q := row-1, col-1
	for p >= 0 && q >= 0 {
		if chessboard[p][q] == "Q" {
			return false
		}
		p--
		q--
	}
	return true
}
