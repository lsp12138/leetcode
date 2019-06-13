package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
