# 53. Maximum Subarray
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:
```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

## 解法一：动态规划

用两个变量，一个result记录最大的和，一个sum记录当前的和。具体：

1. 初始时，result为数组中第一个数的值，sum为0。
2. 然后开始遍历数组：如果sum小于等于0，就让sum等于当前遍历的数。如果sum大于0，就让新的sum等于旧的sum加上当前正遍历的值。新的result等于旧的result和sum的较大值。
3. 最终遍历结束，返回result。

说明：
- 只有sum大于0时，与当前遍历的数加起来才会变大，不然加起来和肯定变小。
- 所以sum小于等于0时，直接把sum变成当前遍历的数就行了。