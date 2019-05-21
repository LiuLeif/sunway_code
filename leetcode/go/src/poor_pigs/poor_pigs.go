// 2018-01-26 17:05
package main

import "fmt"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	ret := 0
	test := (minutesToTest / minutesToDie) + 1
	total := 1
	for total < buckets {
		total *= test
		ret++
	}
	return ret
}
func main() {
	fmt.Println(poorPigs(100, 15, 60))
}
