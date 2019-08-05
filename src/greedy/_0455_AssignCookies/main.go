package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Println(findContentChildren(g, s))
}
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	// 两数组下标
	idxG, idxS := 0, 0
	for idxG < len(g) && idxS < len(s) {
		if g[idxG] <= s[idxS] {
			idxG++
		}
		idxS++
	}
	return idxG
}
