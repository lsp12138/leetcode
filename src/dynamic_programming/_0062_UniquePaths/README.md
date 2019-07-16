# 62. Unique Paths
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

![](https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png)

Note: m and n will be at most 100.

Example 1:
```
Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right
```
Example 2:
```
Input: m = 7, n = 3
Output: 28
```
## 解法一：动态规划
	
不同路径。

采用数组dp[i][j]表示到达（i，j）格子的路径数目，到达 (i,j) 格子的路径数目，等于到达上方格子和左边格子路径数之和：
dp[i][j] = dp[i-1][j] + dp[i][j-1]

另外，到达第一列上任一格子只有1种路径，到达第一行上任一格子只有1种路径。