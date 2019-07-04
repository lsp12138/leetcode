# 105. Construct Binary Tree from Preorder and Inorder Traversal

Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given
```
preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
```
Return the following binary tree:
```
    3
   / \
  9  20
    /  \
   15   7
```
## 解法一：递归

由前序和中序构造二叉树。

前序的第一个结点为根节点，那么中序可以分为左右子树。然后删除前序中的这个根节点（前序剩余部分左部分为左子树，右部分为右子树），对左右子树递归以重复此操作。



