package main

import "fmt"

func main() {
	s := "([)]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		if isOpening(c) {
			stack = append(stack, c)
			continue
		}

		if len(stack) < 1 {
			return false
		}

		lastInd := len(stack) - 1
		lastOpened := stack[lastInd]
		expected := openToClose[lastOpened]
		if c == expected {
			stack = stack[:lastInd]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func isOpening(r rune) bool {
	return r == '(' || r == '[' || r == '{'
}

var openToClose = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}
