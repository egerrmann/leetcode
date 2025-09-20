package main

import "fmt"

func main() {
	s := "abc"
	p := "f*abc"
	fmt.Println(isMatch(s, p))
}

type Effect int

const (
	DefaultEffect Effect = iota
	AnyTimesEffect
)

type Rule struct {
	Symbol byte
	Effect Effect
}

func (r Rule) IsLetterRule() bool {
	return isLetter(r.Symbol)
}

func (r Rule) IsAnyOccurrenceRule() bool {
	return r.Symbol == '.' && r.Effect == AnyTimesEffect
}

func isMatch(s string, p string) bool {
	rules := parsePattern(p)
	anyOccurrenceRuleIndex := -1
	for j, r := range rules {
		if r.IsAnyOccurrenceRule() && anyOccurrenceRuleIndex >= 0 {
			anyOccurrenceRuleIndex = j
		}
		for i := 0; i < len(s); i++ {
			if r.IsLetterRule() && isLetter(s[i])
		}
	}
}

func parsePattern(p string) []Rule {
	var rules []Rule
	for i := 0; i < len(p); i++ {
		rule := Rule{Symbol: p[i]}
		if isLetter(p[i]) {
			if isNextAsterisk(p, i) {
				rule.Effect = AnyTimesEffect
			} else {
				rule.Effect = DefaultEffect
			}
			rules = append(rules, rule)
		} else if p[i] == '.' {
			if isNextAsterisk(p, i) {
				rule.Effect = AnyTimesEffect
			} else {
				rule.Effect = DefaultEffect
			}
		} else if p[i] == '*' {
			continue
		}
	}
	return rules
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z'
}

func isNextAsterisk(p string, currInd int) bool {
	if currInd+1 < len(p) {
		return p[currInd+1] == '*'
	}
	return false
}

//func isMatch(s string, p string) bool {
//	var i, j int
//	var prevChar byte
//	var anyAppearencesMet bool
//	for i, j = 0, 0; i < len(s) && j < len(p); {
//		currChar := s[i]
//		if isLetter(p[j]) {
//			if isNextAsterisk(p, j) {
//				if prevChar != 0 && prevChar != s[i] || s[i] != p[j] {
//					j++
//				}
//				i++
//			} else {
//				if s[i] == p[j] {
//					i, j = i+1, j+1
//				} else {
//					return false
//				}
//			}
//		} else {
//			switch p[j] {
//			case '.':
//				if isNextAsterisk(p, j) {
//					anyAppearencesMet = true
//					if j+2 < len(p) && s[i] == p[j+2] {
//						j = j + 2
//					}
//				}
//				if !isNextAsterisk(p, j) {
//					j++
//				}
//				i++
//			case '*':
//				j++
//			}
//		}
//		prevChar = currChar
//	}
//	if i != len(s) {
//		return false
//	}
//	return true
//}
//
//func isLetter(ch byte) bool {
//	return 'a' <= ch && ch <= 'z'
//}
//
//func isNextAsterisk(p string, currInd int) bool {
//	if currInd+1 < len(p) {
//		return p[currInd+1] == '*'
//	}
//	return false
//}
