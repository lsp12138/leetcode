package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	area := 0
	areaTemp := 0
	for left < right {
		if height[left] < height[right] {
			areaTemp = height[left] * (right - left)
			left++
		} else {
			areaTemp = height[right] * (right - left)
			right--
		}
		if area < areaTemp {
			area = areaTemp
		}
	}
	return area
}
