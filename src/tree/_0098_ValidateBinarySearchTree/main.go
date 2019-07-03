package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 测试用例
}

func isValidBST(root *TreeNode) bool {
	// 分别表示对于当前二叉树来说，值的上下界限，
	// 该二叉树所有的值都应该在它们之间
	min, max := math.MinInt64, math.MaxInt64
	return helper(root, min, max)
}
func helper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	return min < root.Val && root.Val < max &&
		helper(root.Left, min, root.Val) &&
		helper(root.Right, root.Val, max)
}
