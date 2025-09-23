package main

func maxArea(height []int) int {
	var i, j = 0, len(height) - 1
	var res int
	for i != j {
		dist := j - i
		res = max(res, min(height[i], height[j])*dist)
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
