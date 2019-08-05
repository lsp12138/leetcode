package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{1, 2},
		{1, 2},
		{1, 2},
	}
	fmt.Println(eraseOverlapIntervals(intervals))
}

// SortBy 重写排序规则
type SortBy [][]int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i][1] < a[j][1] }

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// sort.Slice(intervals, func(i, j int) bool {
	// 	return intervals[i][1] < intervals[j][1]
	// })
	sort.Sort(SortBy(intervals))
	// 不重叠区间数
	count := 1
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			continue
		}
		end = intervals[i][1]
		count++
	}
	return len(intervals) - count
}
