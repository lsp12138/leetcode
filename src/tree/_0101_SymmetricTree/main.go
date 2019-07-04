package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 测试用例
}

func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

func helper(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && helper(a.Left, b.Right) && helper(a.Right, b.Left)
}
