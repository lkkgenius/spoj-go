package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func mul(split []string) string {
	a := big.NewInt(0)
	b := big.NewInt(0)
	a.SetString(split[0], 10)
	b.SetString(split[1], 10)
	a.Mul(a, b)
	return a.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		fmt.Println(mul(split))
	}
}
