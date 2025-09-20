package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "   -239045092349532099-3"
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	return int(myAtoiInt32(s))
}

func myAtoiInt32(s string) int32 {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if i == 0 && (s[i] == '+' || s[i] == '-') {
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			s = s[:i]
			break
		}
	}
	n, _ := strconv.ParseInt(s, 10, 32)
	return int32(n)
}
