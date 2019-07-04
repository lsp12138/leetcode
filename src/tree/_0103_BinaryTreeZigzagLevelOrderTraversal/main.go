package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 测试用例
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	l := list.New()
	if root == nil {
		return res
	}
	l.PushBack(root)
	// 树的深度
	depth := 0
	for l.Len() != 0 {
		tmp := []int{}
		len := l.Len()
		for i := 0; i < len; i++ {
			f := l.Front()
			node := f.Value.(*TreeNode)
			if depth%2 == 0 {
				tmp = append(tmp, node.Val)
			} else {
				tmp = append([]int{node.Val}, tmp...)
			}
			l.Remove(f)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		depth++
		res = append(res, tmp)
	}
	return res
}
