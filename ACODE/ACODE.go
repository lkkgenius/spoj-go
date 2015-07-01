package main

import (
	"fmt"
)

func acode(in string) int {
	dp := make([]int, len(in))
	dp[0] = 1
	for i := 1; i < len(in); i++ {
		if in[i] != '0' {
			dp[i] = dp[i-1]
		}
		temp := (in[i-1]-'0')*10 + (in[i] - '0')
		if temp >= 10 && temp <= 27 {
			dp[i] += dp[iminus2(i)]
		}
	}
	return dp[len(dp)-1]
}

func iminus2(x int) int {
	if x-2 < 0 {
		return 0
	} else {
		return x - 2
	}
}

func main() {
	var line string
	for {
		fmt.Scanln(&line)
		if line == "0" {
			break
		}
		fmt.Println(acode(line))
	}

}
