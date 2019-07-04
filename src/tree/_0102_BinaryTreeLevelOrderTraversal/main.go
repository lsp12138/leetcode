package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 测试用例
}
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	// 模拟队列
	l := list.New()
	if root == nil {
		return result
	}
	l.PushBack(root)
	for l.Len() != 0 {
		tmp := []int{}
		len := l.Len()
		for i := 0; i < len; i++ {
			f := l.Front()
			// interface{}通过反射转换类型
			node := f.Value.(*TreeNode)
			tmp = append(tmp, node.Val)
			l.Remove(f)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		result = append(result, tmp)
	}
	return result
}
