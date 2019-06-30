# 78. Subsets
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:
```
Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
```
## 解法一：回溯

上去就把单个结果添加到最终结果里头，然后遍历数组各数，更新单个结果，递归，单个结果回退。