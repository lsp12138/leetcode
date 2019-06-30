package main

import "fmt"

func main() {
	m, n := 3, 2
	fmt.Println(uniquePaths(m, n))
}
func uniquePaths(m int, n int) int {
	dp := [100][100]int{}
	// 到达第一列上任一格子只有1种路径
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	// 到达第一行上任一格子只有1种路径
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	// 其余格子路径数等于左边格子和上边格子的和
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
