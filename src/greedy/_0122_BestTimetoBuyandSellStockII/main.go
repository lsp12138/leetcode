package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		dayProfit := prices[i] - prices[i-1]
		if dayProfit > 0 {
			res += dayProfit
		}
	}
	return res
}
