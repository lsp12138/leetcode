# 148. Sort List
Sort a linked list in O(n log n) time using constant space complexity.

Example 1:
```
Input: 4->2->1->3
Output: 1->2->3->4
```
Example 2:
```
Input: -1->5->3->4->0
Output: -1->0->3->4->5
```
## 解法一：归并排序

自底向上的归并排序。

首先，需要用快慢指针法找到链表的中间节点，把原链表拆成两部分，两部分再排序合并成一个链表。这个过程不断递归。