package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(4))
}

func generateParenthesis(n int) []string {
	return foo(0, 0, n, "", []string{})
}

func foo(opensUsed, closesUsed, n int, currStr string, res []string) []string {
	if opensUsed == n && closesUsed == n {
		return append(res, currStr)
	}

	if opensUsed == 0 {
		return foo(opensUsed+1, closesUsed, n, currStr+"(", res)
	} else if opensUsed < n {
		res1 := foo(opensUsed+1, closesUsed, n, currStr+"(", res)
		var res2 []string
		if closesUsed < opensUsed {
			res2 = foo(opensUsed, closesUsed+1, n, currStr+")", res)
		}
		return append(res1, res2...)
	} else {
		return foo(opensUsed, closesUsed+1, n, currStr+")", res)

	}
}
