package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

func combinationSum2(candidates []int, target int) [][]int {
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
		tmp := make([]int, len(resultSingle))
		copy(tmp, resultSingle)
		*result = append(*result, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		// 最终结果需要去重
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		resultSingle = append(resultSingle, candidates[i])
		dfs(candidates, target-candidates[i], i+1, result, resultSingle)
		resultSingle = resultSingle[:len(resultSingle)-1]
	}
}
