// 2018-10-15 13:18
package main

import "fmt"

func numSpecialEquivGroups(A []string) int {
	counter := map[[26]int]bool{}
	for _, s := range A {
		mark := [26]int{}
		for i, c := range s {
			if i%2 == 0 {
				mark[c-'a'] |= 1
			} else {
				mark[c-'a'] |= 2
			}
		}
		counter[mark] = true
	}
	return len(counter)
}
func main() {
	fmt.Println(numSpecialEquivGroups([]string{"aa", "bb", "ab", "ba"}))
}
