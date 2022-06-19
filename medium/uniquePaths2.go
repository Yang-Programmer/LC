package medium

/**
一个机器人位于一个m x n网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。



示例 1：


输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
示例 2：


输入：obstacleGrid = [[0,1],[0,0]]
输出：1


提示：

m ==obstacleGrid.length
n ==obstacleGrid[i].length
1 <= m, n <= 100
obstacleGrid[i][j] 为 0 或 1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/unique-paths-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 放置障碍物的地点说明不可达 dp数组中的值设置为0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	lRow := len(obstacleGrid)
	lCol := len(obstacleGrid[0])
	//如果障碍物在起点或者终点 则没有合适的路径到达终点
	if obstacleGrid[0][0] == 1 || obstacleGrid[lRow-1][lCol-1] == 1 {
		return 0
	}
	dp := make([][]int, lRow)
	for i := 0; i < lRow; i++ {
		dp[i] = make([]int, lCol)
	}
	//初始化dp数组时如果发现有障碍物 dp值为0 且后序也都是0了
	for i := 0; i < lRow && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < lCol && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}
	// 遍历时如果当前节点有障碍物 则该节点dp值为0
	for i := 1; i < lRow; i++ {
		for j := 1; j < lCol; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[lRow-1][lCol-1]
}
