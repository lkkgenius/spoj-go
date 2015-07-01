package main

import (
	"fmt"
	"math/big"
)

func fctrl(x int) *big.Int {
	z := big.NewInt(1)
	for i := 1; i <= x; i++ {
		n := big.NewInt(int64(i))
		z.Mul(z, n)
	}
	return z
}

func main() {
	var r, n int
	fmt.Scan(&r)
	for i := 1; i <= r; i++ {
		fmt.Scan(&n)
		fmt.Println(fctrl(n))
	}
}
