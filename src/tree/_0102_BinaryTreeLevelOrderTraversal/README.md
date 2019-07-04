# 102. Binary Tree Level Order Traversal
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
```
return its level order traversal as:
```
[
  [3],
  [9,20],
  [15,7]
]
```
## 解法一：广度优先（队列）

层序遍历的结果依次压入队列，详见代码。

主要是保证每一次遍历队列时，队列中的值都是一层的。遍历到某节点时，把它的左右子结点加入队列，并从队列中删掉该节点。