# 39. Combination Sum
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

- All numbers (including target) will be positive integers.
- The solution set must not contain duplicate combinations.

Example 1:
```
Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
```
Example 2:
```
Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
```

## 解法一：回溯法

本题所给数组中没有重复的数，且可以重复取数组中的值。先给数组排序。

然后遍历每一个数组中的值（就是把值先放入数组中），本题可以重复取数组中的值所以可以重复遍历，如果各值的和等于target，记录这个结果并撤销数组中最后添加的值（回退），否则直接回退。
