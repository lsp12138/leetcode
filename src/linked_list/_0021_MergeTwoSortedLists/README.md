# 21. Merge Two Sorted Lists
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:
```
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
```
## 解法一：迭代

两个指针：pre指针（指向最终结果的头结点）和cur指针（指向当前结点）。

判断结点大小，调整cur的Next。

## 解法二：递归

还是迭代好！