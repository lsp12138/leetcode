package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 测试用例
}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// last表示反转部分的前一个结点
	// pre作为最后的返回，它的next指向head
	// cur表示当前要反转的结点，它的next指向curNext
	// curPre用于反转
	pre := &ListNode{}
	var last, cur, curNext, curPre *ListNode
	pre.Next = head
	last = pre
	for i := 0; i < m-1; i++ {
		last = last.Next
	}
	// 开始反转
	cur = last.Next
	for i := m; i <= n; i++ {
		curNext = cur.Next
		cur.Next = curPre
		curPre = cur
		cur = curNext
	}
	// 拼接
	last.Next.Next = cur
	last.Next = curPre
	return pre.Next
}
