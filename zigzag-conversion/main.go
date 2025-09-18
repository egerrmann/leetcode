package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcde"
	numRows := 3
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}

	var res strings.Builder

	inc := numRows*2 - 2
	var inc1, inc2 = inc, 0
	var i int
	for base := range inc {
		i = base
		if inc1 < 0 {
			break
		}

		for {
			if inc1 != 0 {
				if i >= len(s) {
					break
				}
				res.WriteByte(s[i])
				i += inc1
			}

			if inc2 != 0 {
				if i >= len(s) {
					break
				}
				res.WriteByte(s[i])
				i += inc2
			}
		}
		inc1 -= 2
		inc2 += 2
	}

	return res.String()
}

// func convert(s string, numRows int) string {
// 	m := initMatrix(numRows, []byte(s))
// 	return collectWord(m)
// }

// func initMatrix(n int, s []byte) [][]*byte {
// 	if n == 1 || n > len(s) {
// 		return convertStrTo2DArr(s)
// 	}

// 	m := (n*2 - 2) * (len(s) / n)

// 	matrix := make([][]*byte, n)
// 	for i := 0; i < n; i++ {
// 		matrix[i] = make([]*byte, m)
// 	}

// 	var i, j int
// 	// zigzagDir true if i is increasing, false otherwise
// 	zigzagDir := true
// 	for _, b := range s {
// 		switch i {
// 		case n - 1:
// 			zigzagDir = false
// 		case 0:
// 			zigzagDir = true
// 		}

// 		matrix[i][j] = &b
// 		if zigzagDir {
// 			i++
// 		} else {
// 			i--
// 			j++
// 		}
// 	}

// 	return matrix
// }

// func convertStrTo2DArr(s []byte) [][]*byte {
// 	res := make([][]*byte, 1)
// 	res[0] = make([]*byte, len(s))
// 	for i := range s {
// 		res[0][i] = &s[i]
// 	}
// 	return res
// }

// func collectWord(matrix [][]*byte) string {
// 	var res strings.Builder
// 	for i := range matrix {
// 		for j := range matrix[i] {
// 			el := matrix[i][j]
// 			if el != nil {
// 				res.WriteByte(*el)
// 			}
// 		}
// 	}

// 	return res.String()
// }
