package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 测试用例
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := root.Val
	helper(root, &maxSum)
	return maxSum
}
func helper(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	// 对该节点的所有孩子递归调用，计算左右子树的最大权值
	leftSum := max(helper(root.Left, maxSum), 0)
	rightSum := max(helper(root.Right, maxSum), 0)
	// 看创建新路径能否获得最大权值
	newSum := root.Val + leftSum + rightSum
	*maxSum = max(*maxSum, newSum)
	// 返回到当前节点的一条最大路径
	return root.Val + max(leftSum, rightSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
