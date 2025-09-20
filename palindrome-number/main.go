package main

import "fmt"

func main() {
	x := 190243
	fmt.Println(isPalindrome(x))
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	var reversed int
	temp := x
	for temp != 0 {
		rem := temp % 10
		reversed = reversed*10 + rem
		temp = temp / 10
	}

	return x == reversed
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	maxInt32Len := 10
	nums := make([]int, 0, maxInt32Len)
	for x != 0 {
		dig := x % 10
		nums = append(nums, dig)
		x /= 10
	}
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}
	return true
}
