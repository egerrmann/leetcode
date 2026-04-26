package main

import (
	"fmt"
	"strings"
)

func main() {
	digits := "223"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	var res []string

	for i := 0; i < len(digits); i++ {
		res = appendLettersByDigit(res, digits[i])
	}

	return res
}

func appendLettersByDigit(combs []string, dig byte) []string {
	digLetters := digToLetter[dig]

	if len(combs) == 0 {
		lettersStr := make([]string, len(digLetters))
		for i, l := range digLetters {
			lettersStr[i] = string(l)
		}
		return lettersStr
	}

	var res []string

	for _, c := range combs {
		for _, dl := range digLetters {
			b := strings.Builder{}
			b.Write(append([]byte(c), dl))
			res = append(res, b.String())
		}
	}

	return res
}

var digToLetter = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}
