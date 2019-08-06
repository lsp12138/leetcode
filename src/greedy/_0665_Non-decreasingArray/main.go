package main

import "fmt"

func main() {
	nums := []int{4, 3, 2}
	fmt.Println(checkPossibility(nums))
}

func checkPossibility(nums []int) bool {
	// 需要修改的元素数目
	count := 0
	for i := 1; i < len(nums) && count < 2; i++ {
		if nums[i-1] <= nums[i] {
			continue
		}
		count++
		if i-2 >= 0 && nums[i-2] > nums[i] {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[i]
		}
	}
	return count < 2
}
