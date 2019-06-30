package main

import "fmt"

func main() {
	n := 3
	fmt.Println(climbStairs(n))
}
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
