package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	// 先排序
	sort.Ints(nums)
	backTrack(nums, &result, []int{}, 0)
	return result
}
func backTrack(nums []int, result *[][]int, resultSingle []int, start int) {
	tmp := make([]int, len(resultSingle))
	copy(tmp, resultSingle)
	*result = append(*result, tmp)
	for i := start; i < len(nums); i++ {
		// 去重
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		resultSingle = append(resultSingle, nums[i])
		backTrack(nums, result, resultSingle, i+1)
		resultSingle = resultSingle[:len(resultSingle)-1]
	}
}
