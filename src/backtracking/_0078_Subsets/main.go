package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
func subsets(nums []int) [][]int {
	result := [][]int{}
	backTrack(nums, &result, []int{}, 0)
	return result
}
func backTrack(nums []int, result *[][]int, resultSingle []int, start int) {
	tmp := make([]int, len(resultSingle))
	copy(tmp, resultSingle)
	*result = append(*result, tmp)
	for i := start; i < len(nums); i++ {
		resultSingle = append(resultSingle, nums[i])
		backTrack(nums, result, resultSingle, i+1)
		resultSingle = resultSingle[:len(resultSingle)-1]
	}
}
