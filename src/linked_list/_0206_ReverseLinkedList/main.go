package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 测试用例
}

func reverseList(head *ListNode) *ListNode {
	// head作为当前要反转的结点
	// pre作为反转后链表的头
	// 反转时让head的next指向pre，pre和head一起后移
	// tmp指向当前要反转结点head的next
	var pre, tmp *ListNode
	for head != nil {
		tmp = head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}
