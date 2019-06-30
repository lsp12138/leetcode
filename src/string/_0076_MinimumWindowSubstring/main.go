package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
func minWindow(s string, t string) string {
	have := map[byte]int{}
	need := map[byte]int{}
	// 所需字母个数
	for i := range t {
		need[t[i]]++
	}
	result := ""
	i, j, count, minLen := 0, 0, 0, len(s)
	for ; j < len(s); j++ {
		// 窗口右端出现了需要的字母，计数器自增
		if have[s[j]] < need[s[j]] {
			count++
		}
		have[s[j]]++
		// 窗口左端有多余的字母，减去
		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}
		// 更新最短串
		curLen := j - i + 1
		if count == len(t) && minLen >= curLen {
			minLen = curLen
			result = s[i : j+1]
		}
	}
	return result
}
