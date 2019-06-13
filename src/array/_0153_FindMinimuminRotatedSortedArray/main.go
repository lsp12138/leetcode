package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
