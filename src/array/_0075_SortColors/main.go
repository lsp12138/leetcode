package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
func sortColors(nums []int) {
	first, second, third := 0, 0, len(nums)-1
	for second <= third {
		switch nums[second] {
		case 0:
			// 这里的second也要右移
			nums[first], nums[second] = nums[second], nums[first]
			first++
			second++
		case 1:
			second++
		case 2:
			nums[second], nums[third] = nums[third], nums[second]
			third--
		}
	}
}
