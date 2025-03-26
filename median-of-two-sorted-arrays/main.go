package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{2, 3}, []int{}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)

	isOdd := (n+m)%2 == 1
	limit := (n + m) / 2

	if n+m == 1 {
		if n == 1 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}

	var i1, i2 int
	prevNum := new(int)
	for i := 0; i < limit; i++ {
		var num1, num2 int
		if i1 == n {
			num1 = math.MaxInt64
		} else {
			num1 = nums1[i1]
		}

		if i2 == m {
			num2 = math.MaxInt64
		} else {
			num2 = nums2[i2]
		}

		if num1 <= num2 {
			i1++
		} else {
			i2++
		}

		if i == limit-1 {
			*prevNum = min(num1, num2)
		}
	}

	var num1, num2 int
	if i1 >= n {
		num1 = math.MaxInt32
	} else {
		num1 = nums1[i1]
	}

	if i2 >= m {
		num2 = math.MaxInt64
	} else {
		num2 = nums2[i2]
	}

	if isOdd {
		return float64(min(num1, num2))
	}

	nextNum := min(num1, num2)
	return (float64(*prevNum) + float64(nextNum)) / 2
}
