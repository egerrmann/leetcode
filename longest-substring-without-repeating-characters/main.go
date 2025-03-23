package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func lengthOfLongestSubstring(s string) int {
	charToInd := make(map[rune]uint16)
	var maxDist uint16
	var ctr uint16
	var minThresholdInd uint16
	for i, ch := range s {
		if ind, ok := charToInd[ch]; ok {
			var currDist uint16
			if minThresholdInd > ind {
				currDist = uint16(i) - minThresholdInd
			} else {
				currDist = uint16(i) - ind
			}
			charToInd[ch] = uint16(i)
			maxDist = max(maxDist, currDist)
			minThresholdInd = max(minThresholdInd, uint16(ind))
			ctr = 0
		} else {
			charToInd[ch] = uint16(i)
			ctr++
			maxDist = max(maxDist, uint16(i)-minThresholdInd, ctr)
		}
	}
	return int(maxDist)
}

// This algo is taken from suggested solutions
func lengthOfLongestSubstringDifferent(s string) int {
	charSet := make(map[rune]struct{})
	var maxDist uint16
	var left uint16
	for right, ch := range s {
		if _, ok := charSet[ch]; ok {
			_, ok1 := charSet[ch]
			for ok1 {
				leftChar := rune(s[left])
				delete(charSet, leftChar)
				left++
				_, ok1 = charSet[ch]
			}

		} else {
			maxDist = max(maxDist, uint16(right)-left+1)
		}
		charSet[ch] = struct{}{}
	}
	return int(maxDist)
}
