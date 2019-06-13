package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	// 先找左区间下标
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		// 从右往左逼近
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	// 没找到相等的数
	if nums[left] != target {
		return result
	}
	// 找到左区间，开始找右区间
	// 注意right的初始取值，考虑在数组[1]中找1的情况
	result[0] = left
	right = len(nums)
	for left < right {
		mid := left + (right-left)/2
		// 从左往右逼近
		// 这个<=很关键，相当于找第一个大于target的数
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	result[1] = left - 1
	return result
}
