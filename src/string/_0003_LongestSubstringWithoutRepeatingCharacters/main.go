package main

import (
	"fmt"
	"math"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	// 字符串转字符数组
	r := []rune(s)
	m := map[rune]int{}
	i, j, maxLen := 0, 0, 0
	for i < len(s) && j < len(s) {
		// 判断滑动窗口右边界的字符在不在set中
		// 不在的话，右边界继续向右划，更新maxlen
		// 在的话，左边界向右划，并在set中剔除左边界字符
		if _, ok := m[r[j]]; !ok {
			m[r[j]] = 0
			j++
			maxLen = int(math.Max(float64(maxLen), float64(j-i)))
		} else {
			delete(m, r[i])
			i++
		}
	}
	return maxLen
}
