# 160. Intersection of Two Linked Lists
Write a program to find the node at which the intersection of two singly linked lists begins.

Notes:

- If the two linked lists have no intersection at all, return null.
- The linked lists must retain their original structure after the function returns.
- You may assume there are no cycles anywhere in the entire linked structure.
- Your code should preferably run in O(n) time and use only O(1) memory.


## 解法一：双指针

如果两个链表相交，那么相交点之后的长度是相同的。所以让两个指针在距离末尾相同的位置开始遍历，遍历到相同值就是相交的点。

所以一个指针肯定在较短的链表的起始位置，我们需要找到较长链表的指针应该从哪里开始。

我们的做法是补全两个链表使之一样长，遍历A链表时我们实际上遍历了A+B链表，遍历B链表时实际上遍历了B+A链表，这样就可以让两个指针在距离末尾相同的位置开始遍历。

具体写代码时，可以让一个链表遍历到末尾时开始从头遍历另一个链表，以达成链表补全的效果。