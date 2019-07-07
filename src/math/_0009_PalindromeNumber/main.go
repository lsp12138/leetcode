package main

import "fmt"

func main() {
	x := 12121
	fmt.Println(isPalindrome(x))
}
func isPalindrome(x int) bool {
	// 负数不是
	if x < 0 {
		return false
	}
	reverse := 0
	tmp := x
	for tmp != 0 {
		reverse = reverse*10 + tmp%10
		tmp /= 10
	}
	return reverse == x
}
