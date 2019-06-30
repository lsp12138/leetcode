package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
func permute(nums []int) [][]int {
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
		visited[i] = true
		resultSingle = append(resultSingle, nums[i])
		backTrack(nums, result, resultSingle, visited)
		visited[i] = false
		resultSingle = resultSingle[:len(resultSingle)-1]
	}
}
