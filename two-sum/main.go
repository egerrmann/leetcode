package main

import (
	"fmt"
)

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	numToInd := map[int32]int32{}
	for i, num := range nums {
		if i == 0 {
			numToInd[int32(num)] = int32(i)
			continue
		}

		diff := target - num
		if ind, ok := numToInd[int32(diff)]; ok {
			return []int{i, int(ind)}
		}

		numToInd[int32(num)] = int32(i)
	}

	return []int{0, 0}
}
