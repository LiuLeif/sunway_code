// 2018-01-30 11:25
package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	splits := make([]string, numRows)
	row := 0
	step := 1
	for i := 0; i < len(s); i++ {
		splits[row] += string(s[i])
		if row == 0 {
			step = 1
		} else if row == numRows-1 {
			step = -1
		}
		row += step
	}
	return strings.Join(splits, "")
}
func main() {
	// P      A       H       N
	//  A   P   L   S   I   I   G
	//    Y       I       R
	fmt.Println(convert("PAYPALISHIRING", 4))
}
