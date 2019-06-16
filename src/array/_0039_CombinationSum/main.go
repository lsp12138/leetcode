package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(nums, target))
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	if len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)
	dfs(candidates, target, 0, &result, []int{})
	return result
}

func dfs(candidates []int, target int, start int, result *[][]int, resultSingle []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*result = append(*result, resultSingle)
		return
	}
	for i := start; i < len(candidates); i++ {
		// 这里需要操作拷贝后的resultSingle，避免多处同时对底层数组进行修改。
		// 这是go语言切片的问题。
		tmp := make([]int, len(resultSingle))
		copy(tmp, resultSingle)
		tmp = append(tmp, candidates[i])
		dfs(candidates, target-candidates[i], i, result, tmp)
		// 回退
		tmp = tmp[:len(tmp)-1]
	}
}
