# 155. Min Stack
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

- push(x) -- Push element x onto stack.
- pop() -- Removes the element on top of the stack.
- top() -- Get the top element.
- getMin() -- Retrieve the minimum element in the stack.
 

Example:
```
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
```
## 解法一：

最小栈。

用两个栈，data和min，再用一个变量记录当前最小值。

data用来正常的push数据，min栈push的时候进行判断只压入最小值。

出栈的时候要记得更新当前最小值。

本题用一个结构体数组模拟栈了，也不用维护这个当前最小值了，因为每次都是和getMin()函数比的。