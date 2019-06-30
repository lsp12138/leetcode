package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 合并后的nums1的指针
	i := m + n - 1
	// 分别指向两个数组最后
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}
	// 如果nums2没遍历完
	for n >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
}
