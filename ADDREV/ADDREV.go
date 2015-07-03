package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(b string) string {
	s := []byte(b)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}
	return string(s)
}

func addrev(x string) int {
	split := strings.Split(x, " ")
	i, _ := strconv.Atoi(reverse(split[0]))
	j, _ := strconv.Atoi(reverse(split[1]))
	result, _ := strconv.Atoi(reverse(strconv.Itoa(i + j)))
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for scanner.Scan() {
		x := scanner.Text()
		fmt.Println(addrev(x))
	}
}
