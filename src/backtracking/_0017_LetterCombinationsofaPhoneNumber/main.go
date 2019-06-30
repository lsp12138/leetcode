package main

import "fmt"

var m = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) != 0 {
		backTrack(digits, &result, "", 0)
	}
	return result
}
func backTrack(digits string, result *[]string, resultSingle string, start int) {
	// 递归结束条件
	if start == len(digits) {
		*result = append(*result, resultSingle)
		return
	}
	// 该数字对应的字符数组
	mapStr := m[string(digits[start])]
	for i := 0; i < len(mapStr); i++ {
		// 处理该组合开头的所有情况
		resultSingle += mapStr[i]
		backTrack(digits, result, resultSingle, start+1)
		// 回退
		resultSingle = resultSingle[:len(resultSingle)-1]
	}
}
