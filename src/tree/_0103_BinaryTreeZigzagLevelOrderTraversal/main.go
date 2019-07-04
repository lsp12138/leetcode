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
	for l.Len() != 0 {
		tmp := []int{}
		len := l.Len()
		for i := 0; i < len; i++ {
			f := l.Front()
			node := f.Value.(*TreeNode)
			if i%2 == 0 {
				tmp = append(tmp, node.Val)
			} else {
				tmp = 
			}
			l.Remove(f)
			
		}
	}
}
