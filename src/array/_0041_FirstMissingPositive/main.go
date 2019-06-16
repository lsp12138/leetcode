package main

import "fmt"

func main() {
	nums := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	for i := range nums {
		// 交换大于等于1的数以及小于等于数组长度的数，放到对应位置，相等就跳过
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	// 返回第一个数值和下标不相等的位置
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	// 如果该数组是个从1开始的顺子
	return len(nums) + 1
}
