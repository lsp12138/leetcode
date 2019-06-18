# 5. Longest Palindromic Substring

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:
```
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```
Example 2:
```
Input: "cbbd"
Output: "bb"
```
## 解法一：动态规划

寻找字符串s的最长回文子串，属于最优子结构问题。

动态规划要找准 “状态” 的定义和 “状态转移方程”。

1. 定义状态
   二维数组dp[i][j]为true表示子串s[i,j]是回文。
2. 状态转义
   如果s[i,j]是回文，那么s[i+1,j-1]也一定是回文，对应的dp[i+1,j-1]也为true。反过来，如果已知dp[i+1,j-1]，就可以通过比较s[i]和s[j]，得到dp[i,j]。

整理可得：dp[i, j] = (s[i] == s[j] and dp[i + 1, j - 1])，并且满足i + 1 <= j - 1。

还要注意，如果s[i]==s[i+1]，那么dp[i][i+1]为true。