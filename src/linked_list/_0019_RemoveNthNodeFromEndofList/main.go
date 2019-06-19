package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 测试用例
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// pre指针
	pre := &ListNode{}
	pre.Next = head
	first, second := pre, pre
	// 先走n+1步
	for i := 0; i <= n; i++ {
		first = first.Next
	}
	// 一起走
	for first != nil {
		first = first.Next
		second = second.Next
	}
	// 删
	second.Next = second.Next.Next
	return pre.Next
}
