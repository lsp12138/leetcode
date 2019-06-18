package main

import (
	"fmt"
	"math"
	"regexp"
)

func main() {
	str := "42"
	fmt.Println(myAtoi(str))
}
func myAtoi(str string) int {
	// 去除开头的空格
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			continue
		}
		str = str[i:]
		break
	}
	reg := regexp.MustCompile(`^[\+\-]?\d+`)
	str = reg.FindString(str)
	// 没匹配到正则
	if str == "" {
		return 0
	}
	result := 0
	switch string(str[0]) {
	case "+":
		result = convert(1, str[1:])
	case "-":
		result = convert(-1, str[1:])
	default:
		result = convert(1, str)
	}
	return result
}

func convert(sign int, str string) int {
	result := 0
	// 用range操作string时，返回下标和Unicode值
	// 表示数字的字符的Unicode值需要减去0字符后再转int
	for _, char := range str {
		result = result*10 + int(char-'0')
		if sign == 1 && result > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && result*-1 < math.MinInt32 {
			return math.MinInt32
		}
	}
	return result * sign
}
