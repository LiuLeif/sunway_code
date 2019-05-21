// 2017-12-27 17:05
package main

import "fmt"

func lastRemaining(n int) int {
	remaining, head, left, step := n, 1, true, 1
	for remaining > 1 {
		if left || remaining%2 == 1 {
			head += step
		}
		step *= 2
		left = !left
		remaining = remaining / 2
	}
	return head
}
func main() {
	// 1 2 3 4 5 6 7 8 9
	//   2   4   6   8
	//   2       6
	//           6
	fmt.Println(lastRemaining(2))
}
