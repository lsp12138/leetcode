package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{
		{10, 16},
		{2, 8},
		{1, 6},
		{7, 12},
	}
	fmt.Println(findMinArrowShots(points))
}

// SortBy 重写排序规则
type SortBy [][]int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i][1] < a[j][1] }

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Sort(SortBy(points))
	// 不重叠区间数
	count := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= end {
			continue
		}
		end = points[i][1]
		count++
	}
	return count
}
