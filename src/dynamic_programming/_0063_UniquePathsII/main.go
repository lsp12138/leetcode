package main

import "fmt"

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := [100][100]int{}
	// 第一列上碰到障碍物，则后边的都过不去
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	// 第一行上碰到障碍物，则后边的都过不去
	for j := 0; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}
	// 其余格子
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
