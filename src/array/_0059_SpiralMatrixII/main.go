package main

import "fmt"

func main() {
	n := 3
	fmt.Println(generateMatrix(n))
}
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	if n == 0 {
		return res
	}
	num := 1
	up, left := 0, 0
	down, right := n-1, n-1
	for up <= down && left <= right {
		// 上部从左到右
		for i := left; i <= right; i++ {
			res[up][i] = num
			num++
		}
		up++
		// 右部从上到下
		for i := up; i <= down; i++ {
			res[i][right] = num
			num++
		}
		right--
		// 下部从右到左
		for i := right; i >= left; i-- {
			res[down][i] = num
			num++
		}
		down--
		// 左部从下到上
		for i := down; i >= up; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}
