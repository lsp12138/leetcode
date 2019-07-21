package main

import "container/list"

func main() {

}

type MyStack struct {
	list *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		list: list.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.list.PushBack(x)
	len := this.list.Len()
	for len > 1 {
		f := this.list.Front()
		this.list.PushBack(f.Value.(int))
		this.list.Remove(f)
		len--
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	f := this.list.Front()
	this.list.Remove(f)
	return f.Value.(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.list.Front().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.list.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
