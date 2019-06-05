package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	// 升序
	sort.Ints(nums)
	result := math.MaxInt64
	// 固定第一个数
	for first := range nums {
		// 跳过重复的数
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 一前一后双指针指向另外两个数
		second := first + 1
		third := len(nums) - 1
		for second < third {
			currentResult := nums[first] + nums[second] + nums[third]
			if math.Abs(float64(target-result)) > math.Abs(float64(target-currentResult)) {
				result = currentResult
			}
			switch {
			case currentResult < target:
				second++
			case currentResult > target:
				third--
			default:
				return target
			}
		}
	}
	return result
}
