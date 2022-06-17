package hard

/**
编写一个程序，通过填充空格来解决数独问题。

数独的解法需 遵循如下规则：

数字1-9在每一行只能出现一次。
数字1-9在每一列只能出现一次。
数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用'.'表示。

board.length == 9
board[i].length == 9
board[i][j] 是一位数字或者 '.'
题目数据 保证 输入数独仅有一个解


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sudoku-solver
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sudoku-solver
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func solveSudoku(board [][]byte) {
	l := len(board)
	var backtracking func(board [][]byte) bool
	backtracking = func(board [][]byte) bool {
		for row := 0; row < l; row++ {
			for col := 0; col < l; col++ {
				if board[row][col] != '.' {
					continue
				}
				for num := '1'; num <= '9'; num++ {
					if isValidSudoNum(row, col, byte(num), board) {
						board[row][col] = byte(num)
						if backtracking(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtracking(board)
}

func isValidSudoNum(row, col int, num byte, board [][]byte) bool {
	//检查同一行是否有元素 num
	l := len(board)
	for i := 0; i < l; i++ {
		if board[row][i] == num {
			return false
		}
	}

	//检查同一列是否有元素 num

	for i := 0; i < l; i++ {
		if board[i][col] == num {
			return false
		}
	}
	// 检查同一 3*3 矩阵是否有相同元素
	startRow := (row / 3) * 3
	strarCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := strarCol; j < strarCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	return true
}
