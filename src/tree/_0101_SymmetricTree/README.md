# 101. Symmetric Tree
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
```
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

But the following [1,2,2,null,3,null,3] is not:
```
    1
   / \
  2   2
   \   \
   3    3
```
## 解法一：深度优先（递归）

如果两个树A和B是对称二叉树，那么满足：

- A与B的根节点值相同
- A的左子树与B的右子树对称
- A的右子树与B的左子树对称

本题传入一个树，可以用辅助函数视为传入两个相同的树。

## 解法二：回文法

如果是对称树，那么层次遍历的结果，每一层都是回文。