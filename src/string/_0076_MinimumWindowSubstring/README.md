# 76. Minimum Window Substring
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:
```
Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
```
Note:

- If there is no such window in S that covers all characters in T, return the empty string "".
- If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

## 解法一：滑动窗口

最小覆盖子串。

参考3题，但更复杂，找出s串的一个最短子串，使他包含t串的所有字母。

在双指针i，j的基础上，还需要两个哈希表，一个have表记录当前窗口中字母的个数，一个need表记录所需要的字母的个数。

计数器count记录当前窗口中，所包含的t中字母的个数，最终他需要和t的长度相等。

具体步骤：首先遍历t串，更新need表。

然后遍历s串更新have表，如果窗口右端出现了需要的字母，计数器自增，如果窗口左端有多余的字母，减去。

同时把当前结果的长度与记录的最短长度比较，更新结果。