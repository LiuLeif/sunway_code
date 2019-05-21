// 2018-01-03 14:30
package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	prev := 0
	count := 1
	curr := 1
	for ; curr < len(chars); curr++ {
		if chars[curr] == chars[prev] {
			count++
		} else {
			if count != 1 {
				copy(chars[prev+1:], []byte(strconv.Itoa(count)))
				prev += len(strconv.Itoa(count))
				count = 1
			}
			prev += 1
			chars[prev] = chars[curr]
		}
	}
	if count != 1 {
		copy(chars[prev+1:], []byte(strconv.Itoa(count)))
		prev += len(strconv.Itoa(count))
	}
	return prev + 1
}
func main() {
	fmt.Println(compress([]byte("aabbcc")))
}
