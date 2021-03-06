package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
func longestPalindrome(s string) string {
	len := len(s)
	if len <= 1 {
		return s
	}
	// 回文起始位置
	start := 0
	// 回文串最大长度
	max := 1
	// 动态规划二维数组
	dp := [][]bool{}
	for i := 0; i < len; i++ {
		dp = append(dp, make([]bool, len))
	}
	// 初始化
	for i := 0; i < len; i++ {
		// 单个字符肯定是回文
		dp[i][i] = true
		// 回文子串长度为2时判断是否回文
		if i < len-1 && s[i] == s[i+1] {
			dp[i][i+1] = true
			max = 2
			start = i
		}
	}
	// l表示回文子串的长度
	for l := 3; l <= len; l++ {
		for i := 0; i+l-1 < len; i++ {
			// 回文终止位置
			j := i + l - 1
			// 状态转移
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				max = l
			}
		}
	}
	return s[start : start+max]
}
