package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-1, 0, -5, -2, -2, -4, 0, 1, -2}
	fmt.Println(fourSum(nums, -9))
}

func fourSum(nums []int, target int) [][]int {
	var res [][]int

	slices.Sort(nums)

	for i, num := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		threeSumRes := threeSum(nums[i+1:], target-num)
		for _, threeSumComb := range threeSumRes {
			res = append(res, append([]int{num}, threeSumComb...))
		}
	}

	return res
}

func threeSum(nums []int, target int) [][]int {
	var res [][]int

	for i := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		j, k := i+1, len(nums)-1

		for j < k {
			a, b, c := nums[i], nums[j], nums[k]

			sum := a + b + c

			if sum > target {
				k--
			} else if sum < target {
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
