# 75. Sort Colors
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:
```
Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```
## 解法一：三指针
	
颜色分类。

一个first指向开头，一个third指向尾部，一个second遍历。

遇到0往前换，first和second右移。

遇到2往后换，third左移。

1就不用管了，second右移。