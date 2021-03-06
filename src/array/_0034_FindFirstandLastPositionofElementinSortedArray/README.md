# 34. Find First and Last Position of Element in Sorted Array

Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:
```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```
Example 2:
```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```
## 解法一：二分
	
在排序数组中查找元素的第一个和最后一个位置。

要求是logn复杂度，用二分。

思路是先找左边界，再找右边界。注意找左边界时由右侧逼近，找右边界时由左侧逼近，反映在代码里就是mid值与target值的大小判断处理有区别。如从右侧逼近就是让right=mid。

