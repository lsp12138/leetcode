# 225. Implement Stack using Queues
Implement the following operations of a stack using queues.

- push(x) -- Push element x onto stack.
- pop() -- Removes the element on top of the stack.
- top() -- Get the top element.
- empty() -- Return whether the stack is empty.
Example:
```
MyStack stack = new MyStack();

stack.push(1);
stack.push(2);  
stack.top();   // returns 2
stack.pop();   // returns 2
stack.empty(); // returns false
```
Notes:

- You must use only standard operations of a queue -- which means only push to back, peek/pop from front, size, and is empty operations are valid.
- Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.
- You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).

## 解法一：

用队列模拟栈。

用一个队列即可。队列push时需要把队尾的元素移动到队首，实现后进先出，方法是队列尾元素之前的全部出队再入队。

1. 栈的push就是队列的push，同时把队尾的元素移动到队首
2. 栈的top就是队列的peek
3. 栈的pop就是队列的remove