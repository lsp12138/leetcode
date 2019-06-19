package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	// 对比数组中剩余的字符串
	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < len(result) && j < len(strs[i]); j++ {
			if result[j] != strs[i][j] {
				break
			}
		}
		result = result[:j]
	}
	return result
}
