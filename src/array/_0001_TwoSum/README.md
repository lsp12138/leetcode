# 1. Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 解法一：暴力遍历

两数之和。

一个循环里套一个循环，计算这两个循环里取到的值的和，是不是等于目标值target。

## 解法二：哈希

第一个循环，把当前值添加到map中作为key，下标作为value（有重复的值即更新下标）。

第二个循环，判断目标值target与当前值的差是不是map的key，且这个差对应的下标不能是当前值的下标（即判断重复值），符合条件返回。

上述两个循环可以合并到一个循环，而且通过调整代码执行顺序可以不用判断重复值。