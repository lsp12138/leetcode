# 4. Median of Two Sorted Arrays
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:
```
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
```
Example 2:
```
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
```
## 解法一：双指针

寻找两个有序数组的中位数。

两个指针i和j分别指向两个排序数组的下标0处，按照数值从小到大的顺序移动指针，并把数值存到切片里，移动次数为两个数组长度和的一半，此时切片最后一个数或者最后两个数的均值就是中位数。

这个方法虽然能提交通过，但是不满足题目要求的logn复杂度。

## 解法二：二分

官方题解看不懂。。。弃了。