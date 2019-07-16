# 9. Palindrome Number
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:
```
Input: 121
Output: true
```
Example 2:
```
Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```
Example 3:
```
Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```
Follow up:

Coud you solve it without converting the integer to a string?

## 解法一：双指针

回文数。

先转成字符串，`s := strconv.Itoa(x)`，然后头尾双指针判断。

## 解法二：反转数字

利用取余和乘以十这两个操作将数字反转，是否和原数字相等。