package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 测试用例
}
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	res := TreeNode{Val: postorder[len(postorder)-1]}
	// 找到中序的根节点的下标
	idx := index(inorder, postorder[len(postorder)-1])
	res.Left = buildTree(inorder[:idx], postorder[:idx])
	res.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
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
