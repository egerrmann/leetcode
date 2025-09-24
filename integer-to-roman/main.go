package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(1994))
}

var romanNums []string = []string{
	"I", "V", "X", "L", "C", "D", "M",
}

func intToRoman(num int) string {
	windowSize := 2
	var res string
	for i := 0; num != 0; i += windowSize {
		tail := num % 10
		switch tail {
		case 0:
		case 4:
			res = fmt.Sprintf("%s%s%s", romanNums[i], romanNums[i+1], res)
		case 9:
			res = fmt.Sprintf("%s%s%s", romanNums[i], romanNums[i+2], res)
		default:
			s1 := romanNums[i]
			var s2 string
			if i+1 < len(romanNums) {
				s2 = romanNums[i+1]
			}
			digitToRoman := digitToRomans(tail, s1, s2)
			res = fmt.Sprintf("%s%s", digitToRoman, res)
		}
		num /= 10
	}
	return res
}

func digitToRomans(num int, s1, s2 string) string {
	var res strings.Builder
	for num != 0 {
		if num >= 5 {
			res.WriteString(s2)
			num -= 5
		} else {
			res.WriteString(s1)
			num -= 1
		}
	}
	return res.String()
}
