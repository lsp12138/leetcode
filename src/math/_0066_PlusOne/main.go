package main

import "fmt"

func main() {
	digits := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))
}
func plusOne(digits []int) []int {
	n := len(digits) - 1
	digits[n]++
	// 加1后开始处理进位
	for i := n; i > 0; i-- {
		// 没有进位直接跳出
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}
	// 处理首位
	if digits[0] >= 10 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}
	return digits
}
