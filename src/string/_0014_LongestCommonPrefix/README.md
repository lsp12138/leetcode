# 14. Longest Common Prefix
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
```
Input: ["flower","flow","flight"]
Output: "fl"
```
Example 2:
```
Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```
Note:

All given inputs are in lowercase letters a-z.

## 解法一：两两对比
	
最长公共前缀。

1. 初始化数组中第一个字符串为初始最长公共前缀。
2. 然后与第二个字符串比较，更新最长公共前缀。
3. 直到最后一个字符串。