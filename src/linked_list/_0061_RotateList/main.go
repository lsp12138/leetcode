package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 测试用例
}

func rotateRight(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	if head == nil || k == 0 {
		return head
	}
	// 简化k的次数
	len := 0
	for i := head; i != nil; i = i.Next {
		len++
	}
	k %= len
	// 快指针先走k步
	for k > 0 {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			fast = head
		}
		k--
	}
	// 一起走直到快指针到尾部
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 此时慢指针就是分割点
	newHead := slow.Next
	fast.Next = head
	slow.Next = nil
	if newHead == nil {
		return head
	}
	return newHead
}
