package main

import "fmt"

func main() {
	s := "LVIII"
	fmt.Println(romanToInt(s))
}
func romanToInt(s string) int {
	m := map[string]int{}
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	// 从右到左遍历
	sum := m[string(s[len(s)-1])]
	for i := len(s) - 2; i >= 0; i-- {
		// 左小右大
		if m[string(s[i])] < m[string(s[i+1])] {
			sum -= m[string(s[i])]
		} else {
			// 左大右小
			sum += m[string(s[i])]
		}
	}
	return sum
}
