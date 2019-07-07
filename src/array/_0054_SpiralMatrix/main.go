package main

import (
	"fmt"
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
	res := []int{}
	if len(matrix) == 0 {
		return res
	}
	up, left := 0, 0
	down, right := len(matrix)-1, len(matrix[0])-1
	for up <= down && left <= right {
		// 上边的行从左到右
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		// 排除只有一行的情况
		if up > down {
			break
		}
		// 右边的列从上到下
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// 排除只有一列的情况
		if right < left {
			break
		}
		// 下边的行从右到左
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		// 左边的列从下到上
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}
