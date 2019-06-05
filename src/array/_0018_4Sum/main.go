package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	// 升序
	sort.Ints(nums)
	// 固定第一个数
	for first := 0; first < len(nums)-3; first++ {
		// 去重
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 固定第二个数
		for second := first + 1; second < len(nums)-2; second++ {
			// 去重，second只和上一个second比较
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 双指针指向两边的两个数
			third, forth := second+1, len(nums)-1
			for third < forth {
				sum := nums[first] + nums[second] + nums[third] + nums[forth]
				switch {
				case sum < target:
					// 小的数需要变大
					third++
				case sum > target:
					// 大的数需要变小
					forth--
				default:
					result = append(result, []int{nums[first], nums[second], nums[third], nums[forth]})
					// 去重
					for third < forth && nums[third] == nums[third+1] {
						third++
					}
					for third < forth && nums[forth] == nums[forth-1] {
						forth--
					}
					third++
					forth--
				}
			}
		}
	}
	return result
}
