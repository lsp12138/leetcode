# 124. Binary Tree Maximum Path Sum
Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

Example 1:
```
Input: [1,2,3]

       1
      / \
     2   3

Output: 6
```
Example 2:
```
Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42
```
## 解法一：深度优先搜索

借助递归辅助函数。具体的：

首先，对该节点的所有孩子递归调用，计算左右子树的最大权值。

其次，看创建新路径能否获得最大权值。创建新路径的权值是当前节点值加上左子树权值和右子树权值，当新路径更好的时候更新最大权值。

最后，返回到当前节点的一条最大路径。