# 142. Linked List Cycle II
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.


Example 1:
```
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

Example 2:
```
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```

Example 3:
```
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
```

Follow-up:
Can you solve it without using extra space?

## 解法一：哈希表

类似141题，不过这题要返回入环的那个节点。

我们可以遍历整个列表并返回哈希表第一个出现重复的节点。

## 解法二：双指针

在141题的基础上，应用一些数学规律，即两指针相遇后，将快指针放到head，然后和慢指针以相同速度一起走，相遇时即入环的节点。