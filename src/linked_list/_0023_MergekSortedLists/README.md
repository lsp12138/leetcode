# 23. Merge k Sorted Lists
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:
```
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
```
## 解法一：分治
	
合并K个排序链表。

21题是合并两个有序链表，本题是合并k个有序链表，采用分治的思想，把本题转化为合并两个有序链表logk次。

首先把k个链表分成两份，每一份都可以再应用一次函数，所以每一份再切分为两份。最后切成一份只有一个链表时应用合并2个链表的函数。