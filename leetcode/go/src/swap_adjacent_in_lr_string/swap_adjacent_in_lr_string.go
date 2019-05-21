// 2018-03-01 11:38
package main

import (
	"fmt"
	"strings"
)

func canTransform(start string, end string) bool {
	if strings.Replace(start, "X", "", -1) != strings.Replace(end, "X", "", -1) {
		return false
	}
	for p1, p2 := 0, 0; p1 < len(start) && p2 < len(end); p1, p2 = p1+1, p2+1 {
		for p1 < len(start) && start[p1] == 'X' {
			p1++
		}
		for p2 < len(end) && end[p2] == 'X' {
			p2++
		}
		if p1 == len(start) && p2 == len(end) {
			return true
		}
		if p1 == len(start) || p2 == len(end) {
			return false
		}
		if start[p1] != end[p2] {
			return false
		}
		if start[p1] == 'L' && p2-p1 > 0 {
			return false
		}
		if start[p1] == 'R' && p1-p2 > 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))
	fmt.Println(canTransform("XXRXXLXXXX", "XXXXRXXLXX"))
	fmt.Println(canTransform("XXXXXXXXLXXXXXXRXXXX", "XXXXXXXXXLXXXXXRXXXX"))
}
