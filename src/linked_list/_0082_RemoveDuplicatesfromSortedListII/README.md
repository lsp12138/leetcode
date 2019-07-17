# 82. Remove Duplicates from Sorted List II
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

Example 1:
```
Input: 1->2->3->3->4->4->5
Output: 1->2->5
```
Example 2:
```
Input: 1->1->1->2->3
Output: 2->3
```
## 解法一：快慢双指针

从有序链表中删除重复节点II。

类似83题，这题是删去重复出现的数，一个也不保留。

快指针跳过重复的数，慢指针拼接。