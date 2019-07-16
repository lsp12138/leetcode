# 11. Container With Most Water

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

![图示](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg)

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example:
```
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
```
## 解法一：暴力法

盛最多水的容器。

一层循环套一层循环。

## 解法二：双指针

用双指针分别指向最前和最后的两个数，然后两个指针相互靠近，靠近时每次移动较小的数对应的指针，因为面积依赖于较小的值。