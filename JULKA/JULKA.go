package main

import (
	"fmt"
	"math/big"
)

func julka(a, b string) (*big.Int, *big.Int) {
	ia, ib := big.NewInt(0), big.NewInt(0)
	ia.SetString(a, 10)
	ib.SetString(b, 10)
	two := big.NewInt(2)
	k, n := big.NewInt(0), big.NewInt(0)
	k.Add(ia, ib)
	k.Div(k, two)
	n.Sub(ia, k)
	return k, n
}

func main() {
	var a, b string
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		k, n := julka(a, b)
		fmt.Println(k)
		fmt.Println(n)
	}
}
