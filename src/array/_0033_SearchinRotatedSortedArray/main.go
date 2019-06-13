package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// 先找最小的数的下标，即旋转点
	// 应用二分时需要比较mid值与right值
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
	minIndex := left
	// 然后让最小的数和target比较，找出target所在区间后应用二分
	left = 0
	right = len(nums) - 1
	if nums[minIndex] <= target && target <= nums[right] {
		left = minIndex
	} else {
		right = minIndex
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
