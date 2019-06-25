package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
func permute(nums []int) [][]int {
	result := [][]int{}
	resultSingle := make([]int, len(nums))
	visited := make([]bool, len(nums))
	backTrack(nums, &result, resultSingle, visited, 0)
	return result
}

func backTrack(nums []int, result *[][]int, resultSingle []int, visited []bool, start int) {

}
