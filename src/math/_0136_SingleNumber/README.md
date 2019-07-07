# 136. Single Number
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:
```
Input: [2,2,1]
Output: 1
```
Example 2:
```
Input: [4,1,2,1,2]
Output: 4
```
# 解法一：哈希表

遍历一遍数组，新出现的数添加到哈希表，再出现就删除，最后哈希表剩下的就是只出现一次的数。

# 解法二：位操作（异或）

概念：

- 如果我们对 0 和二进制位做 XOR 运算，得到的仍然是这个二进制位：
  `a⊕0=a`
- 如果我们对相同的二进制位做 XOR 运算，返回的结果是 0：
  `a⊕a=0`
- XOR 满足交换律和结合律：
  `a⊕b⊕a=(a⊕a)⊕b=0⊕b=b`

所以我们只需要将所有的数进行 XOR 操作，得到那个唯一的数字。
