package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 1, 1, 2}, 1))
}

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)

	closestToTarget := math.MinInt
	minDiffAbs := math.MaxInt

	for i := range nums {
		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			currDiffAbs := diffAbs(sum, target)

			if currDiffAbs == 0 {
				return target
			}

			if currDiffAbs < minDiffAbs {
				closestToTarget = sum
				minDiffAbs = currDiffAbs
			}

			if sum > target {
				k--
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			} else if sum < target {
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			}
		}
	}

	return closestToTarget
}

func diffAbs(a, b int) int {
	diff := a - b
	if diff >= 0 {
		return diff
	} else {
		return -diff
	}
}
