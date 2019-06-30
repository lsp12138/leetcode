package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix))
}
func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 {
		return result
	}
	m := len(matrix)
	n := len(matrix[0])
	// 确认层数
	count := int(math.Min(float64(m), float64(n))+1) / 2
	for i := 0; i < count; i++ {
		// 上边的行从左到右
		for j := i; j < n-i; j++ {
			result = append(result, matrix[i][j])
		}
		// 右边的列从上到下（去除最上边一个）
		for j := i + 1; j < m-i; j++ {
			result = append(result, matrix[j][n-i-1])
		}
		// 下边的行从右到左（去除最右边一个）
		// 还要判断该层是不是只有一行
		for j := n - 1 - i - 1; j >= i && m-1-i != i; j-- {
			result = append(result, matrix[m-1-i][j])
		}
		// 左边的列从下到上（去除最上和最下的）
		// 还要判断该层是不是只有一列
		for j := m - 1 - i - 1; j >= i+1 && n-1-i != i; j-- {
			result = append(result, matrix[j][i])
		}
	}
	return result
}
