package main

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	getTriplets(nums)
}

func getTriplets(nums []int) {

	if len(nums) == 3 && nums[0]+nums[1]+nums[2] == 0 {

	}

	var reviewedTriplets []map[int]int
	var left, right = 0, len(nums) - 1
	for left < right {
		var i, j, k = left, right, -1
		if i+j > 0 {
			k = left + 1
		} else {

		}
	}
}

func areNumbersReviewed(a, b, c int, numToTimesList []map[int]int) bool {
	for _, numToTimes := range numToTimesList {
		if numToTimes[a]+numToTimes[b]+numToTimes[c] == 3 {
			return true
		}
	}
	return false
}
