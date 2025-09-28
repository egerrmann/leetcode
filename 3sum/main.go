package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{0, 0, 0, 0}
	fmt.Println(threeSum(arr))
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	return getTriplets(nums)
}

func getTriplets(nums []int) [][]int {
	if len(nums) == 3 && nums[0]+nums[1]+nums[2] == 0 {
		return [][]int{{nums[0], nums[1], nums[2]}}
	}

	var reviewedTriplets []map[int]int
	for left := 0; left < len(nums)-2; left++ {
		for right := left + 2; right < len(nums); right++ {
			var i, j = left, right
			a, b := nums[i], nums[j]
			k, found := slices.BinarySearch(nums[left+1:right], -1*(a+b))
			if found {
				reviewedTriplets = tryAddNumbersToValid(a, b, nums[left+1+k], reviewedTriplets)
			}
		}
	}

	return collectTriplets(reviewedTriplets)
}

func tryAddNumbersToValid(a, b, c int, numToTimesList []map[int]int) []map[int]int {
	if !areNumbersAdded(a, b, c, numToTimesList) {
		numToTimesList = addNumbersToValid(a, b, c, numToTimesList)
	}
	return numToTimesList
}

func addNumbersToValid(a, b, c int, numToTimesList []map[int]int) []map[int]int {
	numToTimes := make(map[int]int)
	numToTimes[a]++
	numToTimes[b]++
	numToTimes[c]++
	return append(numToTimesList, numToTimes)
}

func areNumbersAdded(a, b, c int, numToTimesList []map[int]int) bool {
	for _, numToTimes := range numToTimesList {
		uniqueNums := make(map[int]struct{})
		uniqueNums[a] = struct{}{}
		uniqueNums[b] = struct{}{}
		uniqueNums[c] = struct{}{}
		totalTimes := 3
		for k := range uniqueNums {
			totalTimes -= numToTimes[k]
		}
		if totalTimes == 0 {
			return true
		}
	}
	return false
}

func collectTriplets(validTriplets []map[int]int) [][]int {
	var triplets [][]int
	for _, numToTimes := range validTriplets {
		arr := make([]int, 0, 3)
		for num, times := range numToTimes {
			for range times {
				arr = append(arr, num)
			}
		}
		triplets = append(triplets, arr)
	}
	return triplets
}
