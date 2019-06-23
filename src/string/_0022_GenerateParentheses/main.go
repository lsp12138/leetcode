package main

import "fmt"

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
func generateParenthesis(n int) []string {
	result := []string{}
	resultSingle := make([]byte, n*2)
	backTrack(n, n, 0, &result, resultSingle)
	return result
}
func backTrack(remainLeft, remainRight, index int, result *[]string, resultSingle []byte) {
	// 递归结束条件
	if remainLeft == 0 && remainRight == 0 {
		*result = append(*result, string(resultSingle))
		return
	}
	// 添加左括号
	if remainLeft > 0 {
		resultSingle[index] = '('
		backTrack(remainLeft-1, remainRight, index+1, result, resultSingle)
	}
	// 添加右括号
	if remainRight > 0 && remainRight > remainLeft {
		resultSingle[index] = ')'
		backTrack(remainLeft, remainRight-1, index+1, result, resultSingle)
	}
}
