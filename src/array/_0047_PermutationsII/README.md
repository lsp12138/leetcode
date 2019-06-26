# 47. Permutations II
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:
```
Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```
## 解法一：回溯

类似46题全排列，不过这题的输入可以有重复的数。所以在46题的基础上去重即可。

去重首先排序，然后在遍历时，如果当前数和上一个数相同，且上一个数没被使用过，就跳过该数。因为如果上一个数没被使用过，后续就会使用到这个和当前一样的数，就会发生重复。