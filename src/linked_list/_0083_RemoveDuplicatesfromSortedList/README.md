# 83. Remove Duplicates from Sorted List
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:
```
Input: 1->1->2
Output: 1->2
```
Example 2:
```
Input: 1->1->2->3->3
Output: 1->2->3
```
## 解法一：遍历即可

从有序链表中删除重复节点。

判断当前结点的和它的next相等不，相等的话让next越过去。