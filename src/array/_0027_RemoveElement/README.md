# 27. Remove Element

Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example 1:
```
Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums being 2.

It doesn't matter what you leave beyond the returned length.
```

Example 2:
```
Given nums = [0,1,2,2,3,0,4,2], val = 2,

Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.

Note that the order of those five elements can be arbitrary.

It doesn't matter what values are set beyond the returned length.
```

## 解法一：双指针

类似26题，一开始两个指针都指向下标0的元素，之后快指针先走，遇到不是val的数就赋值给慢指针。

## 解法二：双指针优化

解法一的缺点是，交换次数过多。如果没有val出现，还是会全部交换一遍。

这时可以一个指针在前，指向左边起第一个出现的val，一个指针在后，指向右边起第一个出现的不是val的数。然后交换即可。