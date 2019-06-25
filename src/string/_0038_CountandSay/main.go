package main

import (
	"fmt"
)

func main() {
	n := 4
	fmt.Println(countAndSay(n))
}
func countAndSay(n int) string {
	buf := []rune{'1'}
	for n > 1 {
		buf = next(buf)
		n--
	}
	return string(buf)

}
func next(buf []rune) []rune {
	// res的长度不会超过buf的2倍
	res := make([]rune, 0, len(buf)*2)
	for i := 0; i < len(buf); {
		count := 1
		for i < len(buf)-1 && buf[i] == buf[i+1] {
			count++
			i++
		}
		res = append(res, rune(count+'0'), buf[i])
		i++
	}
	return res
}
