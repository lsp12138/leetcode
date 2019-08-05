# 435. Non-overlapping Intervals

题目描述：计算让一组区间不重叠所需要移除的区间个数。

```
Input: [ [1,2], [1,2], [1,2] ]

Output: 2

Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.
```

```
Input: [ [1,2], [2,3] ]

Output: 0

Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
```
## 解法一：贪心算法

不重叠的区间个数。

先计算最多能组成的不重叠区间个数，要移除的区间个数就等于区间总个数减去不重叠区间的个数。

思路是先将输入的二维数组按照区间结尾的大小升序排序。遍历区间时，当前区间的起始如果大于前一个区间的结尾，则它们不重叠。
