package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 链表测试用例
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	return mergeTwoLists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	// 一个链表到头了，续上另一个链表
	// 或者一个链表一开始就是空的
	if l1 == nil {
		cur.Next = l2
	}
	if l2 == nil {
		cur.Next = l1
	}
	return pre.Next
}
