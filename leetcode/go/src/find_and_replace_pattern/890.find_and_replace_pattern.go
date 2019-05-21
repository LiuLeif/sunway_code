// 2018-10-25 14:08
package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {
	mark := map[string][]string{}

	for _, word := range words {
		counter := map[rune]rune{}

		curr := 'a'
		newWord := ""
		for _, c := range word {
			if n, ok := counter[c]; ok {
				newWord += string(n)
			} else {
				newWord += string(curr)
				counter[c] = curr
				curr++
			}
		}
		if mark[newWord] == nil {
			mark[newWord] = []string{}
		}
		mark[newWord] = append(mark[newWord], word)
	}

	counter := map[rune]rune{}

	curr := 'a'
	newWord := ""
	for _, c := range pattern {
		if n, ok := counter[c]; ok {
			newWord += string(n)
		} else {
			newWord += string(curr)
			counter[c] = curr
			curr++
		}
	}
	return mark[newWord]
}

func main() {
	fmt.Println(findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "aee"))
}
