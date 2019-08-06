package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
func maxSubArray(nums []int) int {
	result := nums[0]
	sum := 0
	for i := range nums {
		if sum <= 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		result = int(math.Max(float64(sum), float64(result)))
	}
	return result
}
