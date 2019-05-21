// 2017-11-28 18:55
package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	queue := [][2]int{[2]int{0, 0}}
	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true

	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		x, y := top[0], top[1]

		if x+y == len(s3) {
			return true
		}
		if x < len(s1) && s1[x] == s3[x+y] && visited[[2]int{x + 1, y}] == false {
			queue = append(queue, [2]int{x + 1, y})
			visited[[2]int{x + 1, y}] = true
		}
		if y < len(s2) && s2[y] == s3[x+y] && visited[[2]int{x, y + 1}] == false {
			queue = append(queue, [2]int{x, y + 1})
			visited[[2]int{x, y + 1}] = true
		}

	}
	return false
}
func main() {
	// aa  b c      c
	//   db    b ca
	// aadbbcbcac

	// fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabc", "abad", "aabcabad"))
}
