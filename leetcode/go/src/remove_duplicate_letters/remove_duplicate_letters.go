// 2018-02-28 14:20
package main

import "fmt"

func removeDuplicateLetters(s string) string {
	count := map[byte]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	visited := map[byte]bool{}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		count[s[i]]--
		if visited[s[i]] {
			continue
		}
		for len(stack) != 0 && stack[len(stack)-1] > s[i] && count[stack[len(stack)-1]] > 0 {
			visited[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, s[i])
		visited[s[i]] = true
	}
	return string(stack)
}
func main() {
	// acxb
	fmt.Println(removeDuplicateLetters("bacxbc"))
	fmt.Println(removeDuplicateLetters("abcb"))
}
