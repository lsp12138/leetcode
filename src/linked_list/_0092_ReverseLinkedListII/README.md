# 92. Reverse Linked List II
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:
```
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
```
## 解法一：双指针迭代

反转链表II。

类似206题。指定了反转的范围，反转这一部分后和剩下的链表拼接。