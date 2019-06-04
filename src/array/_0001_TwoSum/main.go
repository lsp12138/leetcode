package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for index, value := range nums {
		if i, ok := m[target-value]; ok {
			return []int{i, index}
		}
		m[value] = index
	}
	return nil
}
