package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := map[*ListNode]int{}
	for i := head; i != nil; i = i.Next {
		if _, ok := m[i]; !ok {
			m[i] = 1
		} else {
			return i
		}
	}
	return nil
}
