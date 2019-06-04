package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	// 升序
	sort.Ints(nums)
	result := [][]int{}
	// 固定第一个数
	for first := range nums {
		// 第一个数去重
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 双指针指向另外两个数
		second, third := first+1, len(nums)-1
		for second < third {
			target := nums[first] + nums[second] + nums[third]
			switch {
			case target > 0:
				// 大的数需要变小
				third--
			case target < 0:
				// 小的数需要变大
				second++
			default:
				result = append(result, []int{nums[first], nums[second], nums[third]})
				// 另外两个数去重
				for second < third && nums[second+1] == nums[second] {
					second++
				}
				for second < third && nums[third] == nums[third-1] {
					third--
				}
				second++
				third--
			}
		}
	}
	return result
}
