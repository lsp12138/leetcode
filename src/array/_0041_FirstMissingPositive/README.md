# 41. First Missing Positive
Given an unsorted integer array, find the smallest missing positive integer.

Example 1:
```
Input: [1,2,0]
Output: 3
```
Example 2:
```
Input: [3,4,-1,1]
Output: 2
```
Example 3:
```
Input: [7,8,9,11,12]
Output: 1
```
Note:

Your algorithm should run in O(n) time and uses constant extra space.

## 解法一：下标作为哈希表

遍历一遍数组，把大于等于1的数，以及小于等于数组长度的数，交换放置到数组中对应的位置。如[3,4,-1,1]中的3放置到下标2对应的位置（下标0作为第一个数，所以-1就是第三个数），即和-1交换。

然后再遍历一遍数组，返回第一个数值和下标不相等的位置。