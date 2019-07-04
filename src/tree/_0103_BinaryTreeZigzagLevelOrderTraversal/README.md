# 103. Binary Tree Zigzag Level Order Traversal
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
```
return its zigzag level order traversal as:
```
[
  [3],
  [20,9],
  [15,7]
]
```
## 解法一：广度优先（队列与栈）

类似102题，不过这题是从左到右打印一层，再从右到左打印一层，交替着来。

所以不能只用到队列的先进先出了，反向打印时还需要用到栈的后进先出。

判断先进先出还是后进先出的方式是，根据树的深度，奇数层后进先出（根看做0层）。