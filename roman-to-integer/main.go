package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

var romanToNum = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		s1 := string(s[i])
		var s2 string
		if i+1 < len(s) {
			s2 = string(s[i+1])
		}
		n1, n2 := romanToNum[s1], romanToNum[s2]
		if n2 > n1 {
			res += n2 - n1
			i++
		} else {
			res += n1
		}
	}
	return res
}
