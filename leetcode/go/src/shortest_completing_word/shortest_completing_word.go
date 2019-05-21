// 2018-10-17 09:12
package main

import "fmt"
import "sort"
import "strings"

func shortestCompletingWord(licensePlate string, words []string) string {
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	licensePlate = (strings.ToLower(licensePlate))
	licensePlateCounter := map[rune]int{}
	for _, c := range licensePlate {
		if c >= 97 && c <= 122 {
			licensePlateCounter[c]++
		}
	}
outer:
	for _, s := range words {
		counter := map[rune]int{}
		for _, c := range s {
			counter[c]++
		}
		for k, v := range licensePlateCounter {
			if counter[k] < v {
				continue outer
			}
		}
		return s
	}
	return ""
}
func main() {
	fmt.Println(shortestCompletingWord("GrC8950", []string{"measure", "other", "every", "base", "according", "level", "meeting", "none", "marriage", "rest"}))
}
