# 452. Minimum Number of Arrows to Burst Balloons

题目描述：气球在一个水平数轴上摆放，可以重叠，飞镖垂直投向坐标轴，使得路径上的气球都被刺破。求解最小的投飞镖次数使所有气球都被刺破。

```
Input:
[[10,16], [2,8], [1,6], [7,12]]

Output:
2
```

## 解法一：贪心算法

投飞镖刺破气球。

类似435题。

也是计算不重叠的区间个数，不过和435题的区别在于，[1, 2] 和 [2, 3] 在本题中算是重叠区间。