package main

import (
	"container/list"
)

func main() {

}

type MyQueue struct {
	in  *list.List
	out *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		in:  list.New(),
		out: list.New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.in2out()
	b := this.out.Back()
	this.out.Remove(b)
	return b.Value.(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.in2out()
	b := this.out.Back()
	return b.Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.in.Len() == 0 && this.out.Len() == 0
}

func (this *MyQueue) in2out() {
	if this.out.Len() == 0 {
		for this.in.Len() != 0 {
			b := this.in.Back()
			this.out.PushBack(b.Value.(int))
			this.in.Remove(b)
		}
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
