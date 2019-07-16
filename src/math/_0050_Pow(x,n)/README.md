# 50. Pow(x, n)
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:
```
Input: 2.00000, 10
Output: 1024.00000
```
Example 2:
```
Input: 2.10000, 3
Output: 9.26100
```
Example 3:
```
Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
```
## 解法一：暴力

Pow(x, n)。

超时。

## 解法二：快速幂等法（递归）

使用多次平方进行运算，减少运算次数。

分两种情况，一是偶数次方，2^10=(2^5)^2，2^5就用递归去算。

另一种情况是奇数次方，2^5=(2^2)^2*2。