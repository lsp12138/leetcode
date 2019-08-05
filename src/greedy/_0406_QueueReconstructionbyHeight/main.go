package main

import (
	"sort"
)

func main() {

}

// SortBy 重写排序规则
// 身高 h 降序为主、个数 k 值升序为辅
type SortBy [][]int

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	}
	return a[i][0] > a[j][0]
}

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(SortBy(people))
	res := make([][]int, 0, len(people))
	for i := range people {
		j := people[i][1]
		// 将元素插入到j位置
		res = append(res[:j+1], res[j:]...)
		res[j] = people[i]
	}
	return res
}
