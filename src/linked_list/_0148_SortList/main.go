package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 找到中间结点后，切成两段递归
	mid := findMid(head)
	right := sortList(mid.Next)
	mid.Next = nil
	left := sortList(head)
	return merge(left, right)
}

// 快慢指针法找到中间节点
func findMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 对两组链表进行排序合并
func merge(left, right *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	for left != nil && right != nil {
		if left.Val > right.Val {
			cur.Next = right
			right = right.Next
		} else {
			cur.Next = left
			left = left.Next
		}
		cur = cur.Next
	}
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}
	return pre.Next
}
