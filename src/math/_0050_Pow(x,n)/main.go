package main

import "fmt"

func main() {
	x := 2.00000
	n := 10
	fmt.Println(myPow(x, n))
}
func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return pow(x, n)
}
func pow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	half := pow(x, n/2)
	if n%2 == 1 {
		return half * half * x
	}
	return half * half
}
