package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	return int(reverseInt32(int32(x)))
}

func reverseInt32(x int32) int32 {
	var result int32
	for x != 0 {
		tail := x % 10
		result = result*int32(10) + int32(tail)
		if result-result/10*10 != tail {
			return 0
		}
		x /= 10
	}
	return result
}
