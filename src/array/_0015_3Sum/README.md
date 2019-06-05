# 15. 3Sum

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:
```
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

## 解法一：暴力循环

三重循环可解之。考虑到结果不能重复，可以先排序，然后判断时跳过重复值。

肯定超时。

## 解法二：双指针

还是先排序，然后固定一个数，找另外两个值，使他们的和为0。

如先固定第一个数，然后一个指针指向第二个数，一个指针指向最后一个数，两个指针相互靠近时跳过重复值。