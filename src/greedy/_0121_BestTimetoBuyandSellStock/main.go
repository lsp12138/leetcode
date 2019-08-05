package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
func maxProfit(prices []int) int {
	// 最后的收益，临时收益
	res, tmpRes := 0, 0
	for i := 1; i < len(prices); i++ {
		tmpRes += prices[i] - prices[i-1]
		if tmpRes < 0 {
			tmpRes = 0
		}
		if res < tmpRes {
			res = tmpRes
		}
	}
	return res
}
