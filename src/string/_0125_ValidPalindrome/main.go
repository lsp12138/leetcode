package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	s = strings.ToLower(s)
	front := 0
	back := len(s) - 1
	for front < back {
		// 跳过不是字母和数字的字符
		for front < back && !helper(s[front]) {
			front++
		}
		for front < back && !helper(s[back]) {
			back--
		}
		if s[front] != s[back] {
			return false
		}
		front++
		back--
	}
	return true
}
func helper(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
