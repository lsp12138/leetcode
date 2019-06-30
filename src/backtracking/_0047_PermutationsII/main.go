package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
func permuteUnique(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	result := [][]int{}
	resultSingle := []int{}
	visited := make([]bool, len(nums))
	backTrack(nums, &result, resultSingle, visited)
	return result
}
func backTrack(nums []int, result *[][]int, resultSingle []int, visited []bool) {
	if len(resultSingle) == len(nums) {
		tmp := make([]int, len(resultSingle))
		copy(tmp, resultSingle)
		*result = append(*result, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == true {
			continue
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] == false {
			continue
		}
		visited[i] = true
		resultSingle = append(resultSingle, nums[i])
		backTrack(nums, result, resultSingle, visited)
		visited[i] = false
		resultSingle = resultSingle[:len(resultSingle)-1]
	}
}
