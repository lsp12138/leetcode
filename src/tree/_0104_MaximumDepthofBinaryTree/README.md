# 104. Maximum Depth of Binary Tree
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
```
## 解法一：深度遍历
	
二叉树的最大深度。

树的深度等于左右子树的深度的较大值加上1，用递归。