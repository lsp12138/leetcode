# 88. Merge Sorted Array
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

- The number of elements initialized in nums1 and nums2 are m and n respectively.
- You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:
```
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
```
## 解法一：双指针（从后往前）

合并两个有序数组。

把排好序的nums1的前m项和nums2的前n项合并，放入nums1中。

这里需要三个指针，m和n分别指向两个数组最后的数，i指向合并后的nums1的最后。

遍历条件是m和n都大于等于0，大的数放i的位置。

遍历完毕后还要判断nums2有没有遍历完，如果没有，把nums2剩余的数移动到nums1。