package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 测试用例
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	res := TreeNode{Val: preorder[0]}
	// 找到中序数组中的根节点下标
	idx := index(inorder, preorder[0])
	// 两个数组长度要相同，先序数组被一分为二
	res.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	res.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return &res
}

func index(array []int, target int) int {
	for idx := range array {
		if array[idx] == target {
			return idx
		}
	}
	return -1
}
