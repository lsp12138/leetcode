package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}

// 从前先后
// func canJump(nums []int) bool {
// 	start, end := 0, 0
// 	for start <= end && end < len(nums)-1 {
// 		end = int(math.Max(float64(end), float64(nums[start]+start)))
// 		start++
// 	}
// 	return end >= len(nums)-1
// }

// 从后向前
func canJump(nums []int) bool {
	start, end := len(nums)-2, len(nums)-1
	for start >= 0 {
		if nums[start]+start >= end {
			end = start
		}
		start--
	}
	return end <= 0
}
