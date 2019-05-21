// 2018-01-11 16:31
package main

import "fmt"

func reverse(s string) string {
	t := []byte(s)
	l := len(t)
	for i := 0; i < l/2; i++ {
		t[i], t[l-i-1] = t[l-i-1], t[i]
	}
	return string(t)
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func palindromePairs(words []string) [][]int {
	ret := [][]int{}
	m := map[string]int{}
	for i := 0; i < len(words); i++ {
		m[words[i]] = i
	}
	visited := map[[2]int]bool{}
	for i := 0; i < len(words); i++ {
		word := words[i]
		for j := 0; j <= len(word); j++ {
			str1 := word[:j]
			str2 := word[j:]
			if index, ok := m[reverse(str2)]; index != i && isPalindrome(str1) && ok {
				if !visited[[2]int{index, i}] {
					ret = append(ret, []int{index, i})
				}
				visited[[2]int{index, i}] = true
			}
			if index, ok := m[reverse(str1)]; index != i && isPalindrome(str2) && ok {
				if !visited[[2]int{i, index}] {
					ret = append(ret, []int{i, index})
				}
				visited[[2]int{i, index}] = true
			}
		}
	}
	return ret
}
func main() {
	fmt.Println(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll", "a"}))
	fmt.Println(palindromePairs([]string{"ab", "ba"}))
	fmt.Println(palindromePairs([]string{"aba", ""}))
}
