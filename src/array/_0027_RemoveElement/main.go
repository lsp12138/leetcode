package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}

// func removeElement(nums []int, val int) int {
// 	// 解法一：快慢双指针（一开始俩指针是重叠的）
// 	// 只要快指针的值不是val，就赋给慢指针。
// 	slow := 0
// 	for fast := 0; fast < len(nums); fast++ {
// 		if nums[fast] != val {
// 			nums[slow] = nums[fast]
// 			slow++
// 		}
// 	}
// 	return slow
// }

func removeElement(nums []int, val int) int {
	// 解法二：前后双指针
	// 前指针指向左边起第一个val，后指针指向右边起第一个不是val的数
	// 然后前后交换
	left := 0
	right := len(nums) - 1
	for {
		for left < len(nums) && nums[left] != val {
			left++
		}
		for right >= 0 && nums[right] == val {
			right--
		}
		if left >= right {
			break
		}
		// go语言特有的元素交换写法
		nums[left], nums[right] = nums[right], nums[left]
	}
	return left
}
