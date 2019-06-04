package main

import (
	"sort"
	"fmt"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	
}