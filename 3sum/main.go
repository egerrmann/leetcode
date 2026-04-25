package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-2, 0, 1, 1, 2}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var res [][]int

	slices.Sort(nums)

	for i := range nums {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		j, k := i+1, len(nums)-1

		for j < k {
			a, b, c := nums[i], nums[j], nums[k]
			if a+b > 0 {
				break
			}

			sum := a + b + c

			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				res = append(res, []int{a, b, c})
				j++
				k--

				for nums[j] == nums[j-1] && j < k {
					j++
				}

				for nums[k] == nums[k+1] && j < k {
					k--
				}
			}
		}
	}

	return res
}
