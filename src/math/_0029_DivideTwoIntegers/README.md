# 29. Divide Two Integers
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:
```
Input: dividend = 10, divisor = 3
Output: 3
```
Example 2:
```
Input: dividend = 7, divisor = -3
Output: -2
```
Note:

- Both dividend and divisor will be 32-bit signed integers.
- The divisor will never be 0.
- Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.

## 解法一：用减法实现除法

两数相除。

不能使用乘法、除法和取余运算，编写一个除法函数。

可以采用减法实现除法：被除数减去除数，减一次计数器加一，直到被除数小于0，此时计数器结果就是答案。

但这样做遇到很大的被除数和很小的除数时会超时，需要改良。

改良方法是每次做减法时除数加倍，对应的计数器也加倍，请看表格：

|次数|被除数|除数|计数器|
|---|---|---|---|
|初始|15|3|0|
|1|12|6|0+1=1|
|2|6|12|1+1x2=3|

此时除数大于被除数，重新循环：

|次数|被除数|除数|计数器|
|---|---|---|---|
|初始|6|3|3|
|1|3|6|3+1=4|

同理，再次重新开始：

|次数|被除数|除数|计数器|
|---|---|---|---|
|初始|3|3|4|
|1|0|6|4+1=5|

被除数为0，结束。计数器为5即答案。

因为乘法、除法、取余计算都不可以使用，所以用位运算替代。计算之前需要用异或运算提取符号，具体看代码。