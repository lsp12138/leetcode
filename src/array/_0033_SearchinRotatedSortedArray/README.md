# 33. Search in Rotated Sorted Array

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:
```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```
Example 2:
```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

## 解法一：二分查找

一个数组内分为两个区间分别升序，且无重复数，然后要求查找一个值，且复杂度为logn，自然想到二分查找。

我们首先找到两个区间的交点，即最小的数。方法是二分时比较mid值与right值的大小，如果mid值大于right值，最小的值就在mid值右侧。

然后判断target在哪个区间，在对应区间应用二分查找即可。