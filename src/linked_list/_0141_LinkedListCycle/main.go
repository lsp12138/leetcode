package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		// 没环的判断
		if fast == nil || fast.Next == nil {
			return false
		}
		// 快指针走2步
		slow, fast = slow.Next, fast.Next.Next
	}
	return true
}
