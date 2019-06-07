package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	// 快慢指针
	slow := 0
	fast := slow + 1
	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast++
		} else {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	// 返回无重复值的这部分的长度
	return slow + 1
}
