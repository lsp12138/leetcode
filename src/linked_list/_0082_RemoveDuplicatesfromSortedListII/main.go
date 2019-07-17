package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 测试用例
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := &ListNode{}
	pre.Next = head
	slow, fast := pre, head
	for fast != nil {
		// 跳过重复值
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}
		// 是否拼接
		if slow.Next == fast {
			slow = slow.Next
			fast = fast.Next
		} else {
			slow.Next = fast.Next
			fast = fast.Next
		}
	}
	return pre.Next
}
