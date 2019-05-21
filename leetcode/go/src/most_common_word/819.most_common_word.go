// 2018-11-01 09:25
package main

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	isAscii := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
	}
	counter := map[string]int{}
	prev := -1
	for i := 0; i < len(paragraph); i++ {
		if !isAscii(paragraph[i]) {
			if prev != -1 {
				counter[strings.ToLower(paragraph[prev:i])]++
				prev = -1
			}
		} else {
			if prev == -1 {
				prev = i
			}
		}
	}
	if prev != -1 {
		counter[strings.ToLower(paragraph[prev:])]++
	}

	for _, v := range banned {
		delete(counter, v)
	}

	maxCount := 0
	maxValue := ""
	for k, v := range counter {
		if v >= maxCount {
			maxCount = v
			maxValue = k
		}
	}
	return maxValue
}

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}
