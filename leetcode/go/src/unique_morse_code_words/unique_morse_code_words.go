// 2018-04-13 10:53
package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	m := map[string]bool{}
	table := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	for _, word := range words {
		morse := ""
		for _, c := range word {
			morse += table[c-'a']
		}
		m[morse] = true
	}
	return len(m)
}
func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
