package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	m, n := len(nums1), len(nums2)
	mid := (m + n) / 2
	tmp := []int{}
	var result float64
	for i+j <= mid {
		// 要考虑其中一个数组遍历完了的情况
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				tmp = append(tmp, nums1[i])
				i++
			} else {
				tmp = append(tmp, nums2[j])
				j++
			}
		} else if i < m && j >= n {
			tmp = append(tmp, nums1[i])
			i++
		} else {
			tmp = append(tmp, nums2[j])
			j++
		}
	}
	if (m+n)%2 == 1 {
		result = float64(tmp[mid])
	} else {
		result = (float64(tmp[mid] + tmp[mid-1])) * 0.5
	}
	return result
}
