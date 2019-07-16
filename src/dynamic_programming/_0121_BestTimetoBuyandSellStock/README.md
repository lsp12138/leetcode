# 121. Best Time to Buy and Sell Stock
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:
```
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
```
Example 2:
```
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```
## 解法一：动态规划
	
买卖股票的最佳时机。

需要两个变量，一个记录最终的收益，一个记录临时的收益。

遍历一遍输入，每次都更新临时收益，使临时收益=临时收益+今天的收益。如果临时收益小于0，就重新赋值为0。取最大的临时收益为最终收益。