package main

import "fmt"

func main() {
	numRows := 5
	fmt.Println(generate(numRows))
}
func generate(numRows int) [][]int {
	// 前两行
	result := [][]int{[]int{1}, []int{1, 1}}
	// 从第三行继续生成杨辉三角
	for i := 2; i < numRows; i++ {
		// 第一个数总是1
		tmp := []int{1}
		for j := 1; j < i; j++ {
			tmp = append(tmp, result[i-1][j-1]+result[i-1][j])
		}
		// 最后一个数也是1
		tmp = append(tmp, 1)
		result = append(result, tmp)
	}
	// 返回指定行数
	return result[:numRows]
}
