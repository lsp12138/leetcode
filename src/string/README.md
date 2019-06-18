# 字符串

## 字符串处理相关

1. 对于字符串s来说，直接用s[1]取出来的是byte类型。它与其它string类型比较时需要string(s[1])转化一下。
2. 用range操作string时，返回下标和Unicode值。表示数字的字符需要转int时，要把Unicode值减去0字符'0'后再转int。参考8题。

## 滑动窗口

参考3。