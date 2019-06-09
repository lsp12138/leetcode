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
	rotateIndex := findRotateIndex(nums)
	left := 0
	right := len(nums) - 1
	if nums[rotateIndex] == target {
		return rotateIndex
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	if nums[left] > target {
		return binarySearch(nums, rotateIndex, right, target)
	}
	return binarySearch(nums, left, rotateIndex, target)

}

// 找到旋转数组中最小的数的下标
func findRotateIndex(nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2
	for left < right {
		if nums[mid] > nums[mid+1] {
			return mid + 1
		}
		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}

// 二分查找
func binarySearch(nums []int, left int, right int, target int) int {
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
