# 20. Valid Parentheses
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Example 1:
```
Input: "()"
Output: true
```
Example 2:
```
Input: "()[]{}"
Output: true
```
Example 3:
```
Input: "(]"
Output: false
```
Example 4:
```
Input: "([)]"
Output: false
```
Example 5:
```
Input: "{[]}"
Output: true
```
## 解法一：栈

有效的括号。

遇到左括号压入，遇到右括号时弹出压入的左括号。最后栈为空则满足条件。

可惜go语言没有栈，用数组加移动的下标模拟，入栈下标右移，出栈下标左移，最后下标为0则满足条件。