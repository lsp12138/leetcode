# 7. Reverse Integer
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
```
Input: 123
Output: 321
```
Example 2:
```
Input: -123
Output: -321
```
Example 3:
```
Input: 120
Output: 21
```
## 解法一：逐位拆解

用除和取余直接拆，拆完再加起来。

注意，这题指定32位，反转后需要判断是否溢出。