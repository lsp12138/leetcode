package main

import "fmt"

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
func isValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]rune, len(s))
	index := 0
	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack[index] = m[char]
			index++
		case ')', ']', '}':
			if index > 0 && stack[index-1] == char {
				index--
			} else {
				return false
			}
		}
	}
	return index == 0
}
