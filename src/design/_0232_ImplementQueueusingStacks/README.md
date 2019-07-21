# 232. Implement Queue using Stacks
Implement the following operations of a queue using stacks.

- push(x) -- Push element x to the back of queue.
- pop() -- Removes the element from in front of queue.
- peek() -- Get the front element.
- empty() -- Return whether the queue is empty.
Example:
```
MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);  
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false
```
Notes:

- You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
- Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
- You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).

## 解法一：

用栈实现队列。

两个栈，一个in一个out，值从in进来再到out再出来，两次后进先出就成先进先出了。

1. 队列的push操作就是in的push。
2. 队列的peek操作就是out的peek。
3. 队列的pop操作就是out的pop。

需要注意的是，队列的peek和pop之前要保证out非空。