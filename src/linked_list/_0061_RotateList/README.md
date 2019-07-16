# 61. Rotate List
Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:
```
Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
```
Example 2:
```
Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL
```
## 解法一：双指针

旋转链表。

快慢双指针，找到要分割的点，找时不能一个个遍历k遍，会超时，需要用取余来处理，使k=k%链表长度，这样可以简化遍历次数。

具体的，快指针先走k步，然后一起走直到快指针到尾部，此时慢指针就是分割点。