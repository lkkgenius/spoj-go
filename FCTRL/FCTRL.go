package main

import (
	"fmt"
)

func Z(x int) int {
	ret := 0
	for i := 5; i <= x; i *= 5 {
		ret += x / i
	}
	return ret
}
func main() {
	var r, x int
	fmt.Scan(&r)
	for i := 1; i <= r; i++ {
		fmt.Scan(&x)
		fmt.Println(Z(x))
	}
}
