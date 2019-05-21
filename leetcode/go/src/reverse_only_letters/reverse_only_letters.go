// 2018-10-12 15:15
package main

import "fmt"

func reverseOnlyLetters(S string) string {
	v := []rune(S)
	i, j := 0, len(v)-1

	isAscii := func(c rune) bool {
		if c >= 97 && c <= 122 {
			return true
		}
		if c >= 65 && c <= 90 {
			return true
		}
		return false
	}

	for i < j {
		if !isAscii(v[i]) {
			i++
			continue
		}
		if !isAscii(v[j]) {
			j--
			continue
		}
		v[i], v[j] = v[j], v[i]
		i++
		j--
	}
	return string(v)

}
func main() {
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}
