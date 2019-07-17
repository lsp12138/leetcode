# 19. Remove Nth Node From End of List
Given a linked list, remove the n-th node from the end of list and return its head.

Example:
```
Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
```
Note:

Given n will always be valid.

## 解法一：双指针

删除链表的倒数第N个节点。

先设置一个pre指针，它指向头结点。双指针一开始都在pre的位置。

指针A先走n+1步，然后另一个指针B开始一起走，直到A走到头变为空指针。此时，指针B的Next元素就是该被删除的元素，让B.Next=B.Next.Next即可。