# 55. Jump Game
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:
```
Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```
xample 2:
```
Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
```
## 解法一：贪心

从前向后遍历，两个指针，一个start指针用来遍历，另一个end指针记录当前最远能走到的位置下标，遍历结束后如果这个下标能够到达最后一个元素，返回true。

因为要判断能不能到达最后一个元素，所以在遍历时，指针最远只能到倒数第二个元素。

## 解法二：贪心

从后向前遍历，两个指针，end指针指向最后一个元素，start指针指向倒数第二个。判断start位置能不能到达end位置，能的话end指向start的位置，即问题转化为能否到达新的end位置。

遍历结束后看end到没到第一个元素的位置。