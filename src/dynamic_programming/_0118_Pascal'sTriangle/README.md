# 118. Pascal's Triangle
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.

![](https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif)

In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:
```
Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
```
## 解法一：动态规划

杨辉三角。某行的值由上一行相邻的两数相加得到。