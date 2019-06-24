package main

import (
	"fmt"
	"math"
)

func main() {
	dividend, divisor := 15, 3
	fmt.Println(divide(dividend, divisor))
}
func divide(dividend int, divisor int) int {
	result := 0
	// 提取符号，-1代表负数，1代表正数
	sign := 1
	// 异或计算是不同则为1，符号位为1表示为负，
	// 所以符号位异或之后，如果为1，则原值一正一负，
	// 此时异或的结果为负。
	if dividend^divisor < 0 {
		sign = -1
	}
	// 提取符号后取绝对值
	if dividend < 0 {
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisor = 0 - divisor
	}
	for dividend >= divisor {
		count := 1
		curDivisor := divisor
		for dividend >= curDivisor {
			dividend -= curDivisor
			result += count
			// 左移1位表示乘以2
			count <<= 1
			curDivisor <<= 1
		}
	}
	// 恢复符号
	if sign == -1 {
		result = 0 - result
	}
	// 越界判断
	if result > math.MaxInt32 || result < math.MinInt32 {
		return math.MaxInt32
	}
	return result
}
