# 605. Can Place Flowers

Suppose you have a long flowerbed in which some of the plots are planted and some are not. However, flowers cannot be planted in adjacent plots - they would compete for water and both would die.

Given a flowerbed (represented as an array containing 0 and 1, where 0 means empty and 1 means not empty), and a number n, return if n new flowers can be planted in it without violating the no-adjacent-flowers rule.
```
Example 1:
Input: flowerbed = [1,0,0,0,1], n = 1
Output: True
```
```
Example 2:
Input: flowerbed = [1,0,0,0,1], n = 2
Output: False
```
## 解法一：一次遍历（贪心算法？）

种植花朵。

还能否种下n朵花。一次遍历，如果当前值是0，判断它前一个数和后一个数是不是0。