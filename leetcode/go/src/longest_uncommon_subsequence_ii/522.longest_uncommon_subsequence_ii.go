// 2018-10-31 17:06
package main

import (
	"fmt"
	"sort"
)

func findLUSlength(strs []string) int {
	contains := func(str, substr string) bool {
		// abcde
		// bde
		j := 0
		for i := 0; i < len(str) && j != len(substr); i++ {
			if str[i] == substr[j] {
				j++
			}
		}
		return j == len(substr)
	}

	counter := map[string]int{}
	for _, s := range strs {
		counter[s]++
	}

	uniqStrs := []string{}
	for k, _ := range counter {
		uniqStrs = append(uniqStrs, k)
	}

	sort.Slice(uniqStrs, func(i, j int) bool {
		return len(uniqStrs[i]) > len(uniqStrs[j])
	})
loop:
	for i := 0; i < len(uniqStrs); i++ {
		if counter[uniqStrs[i]] != 1 {
			continue
		}
		for j := 0; j < i; j++ {
			if contains(uniqStrs[j], uniqStrs[i]) {
				continue loop
			}
		}
		return len(uniqStrs[i])
	}
	return -1
}

func main() {
	fmt.Println(findLUSlength([]string{"aabbcc", "aabbcc", "cb", "abc"}))
}
