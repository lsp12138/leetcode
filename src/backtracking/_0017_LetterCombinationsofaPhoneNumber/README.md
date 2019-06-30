# 17. Letter Combinations of a Phone Number

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

![图](http://upload.wikimedia.org/wikipedia/commons/thumb/7/73/Telephone-keypad2.svg/200px-Telephone-keypad2.svg.png)

Example:
```
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
```
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.

## 解法一：回溯

用哈希表保存数字与字母的对应关系。类似排列组合问题。

除了原始输入，递归函数需要3个参数：总结果，单个结果，起始下标。

代码大致顺序：

1. 递归结束条件（下标越界）
2. 更新单个结果，然后调用递归函数
3. 递归函数调用完毕后注意回退（主要是单个结果的回退）


