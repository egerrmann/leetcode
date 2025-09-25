package main

import "fmt"

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	longestPrefix := strs[0]
	if longestPrefix == "" {
		return ""
	}

	for i := 1; i < len(strs); i++ {
		if strs[i] == "" {
			return ""
		}
		if len(strs[i]) < len(longestPrefix) {
			longestPrefix = longestPrefix[:len(strs[i])]
		}
		for j := 0; j < len(longestPrefix); j++ {
			if longestPrefix[j] != strs[i][j] {
				longestPrefix = longestPrefix[:j]
				break
			}
		}
		if longestPrefix == "" {
			return ""
		}
	}
	return longestPrefix
}
