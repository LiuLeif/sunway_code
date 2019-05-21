// 2018-01-09 15:53
package main

import "fmt"

func reverseVowels(s string) string {
	vowels := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels = append(vowels, s[i])
		}
	}
	ret := ""
	vowelsIndex := 1
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			ret = ret + string(vowels[len(vowels)-vowelsIndex])
			vowelsIndex++
		default:
			ret = ret + string(s[i])
		}
	}
	return ret
}

func isVowels(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func reverseVowelsTwoPointer(s string) string {
	lo, hi := 0, len(s)-1
	t := []byte(s)
	for lo < hi {
		for lo < hi && !isVowels(s[lo]) {
			lo++
		}
		for lo < hi && !isVowels(s[hi]) {
			hi--
		}
		t[lo], t[hi] = t[hi], t[lo]
		lo++
		hi--
	}

	return string(t)
}

func main() {
	fmt.Println(reverseVowelsTwoPointer("hello"))
}
