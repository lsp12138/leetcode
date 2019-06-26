# 49. Group Anagrams
Given an array of strings, group anagrams together.

Example:
```
Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
```
Note:

- All inputs will be in lowercase.
- The order of your output does not matter.

## 解法一：哈希表

两个单词中的各个字母出现的次数相同，则为字母异位词。

我们遍历字符串时，先排序，然后把排好序的字符串作为map的key，value即为各个符合条件的单词。

go的字符串排序【sort.Strings()】真是恶心，一个字符串要先转成数组，排序完再转成字符串。