// 2017-11-14 18:54
package main

import "fmt"

func ReverseString(s string) string {
	data := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		data[i], data[len(s)-i-1] = data[len(s)-i-1], data[i]
	}
	return string(data)
}

func IsParlindrome(s string) bool {
	return s == ReverseString(s)
}

func longestPalindrome(s string) string {
	rev := ReverseString(s)
	max_index := 0
	max_len := 0
	m := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		m[i] = make([]int, len(s)+1)
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			if s[i-1] == rev[j-1] {
				m[i][j] = m[i-1][j-1] + 1
				if max_len < m[i][j] && IsParlindrome(string(s[i-m[i][j]:i])) {
					max_len = m[i][j]
					max_index = i
				}

			} else {
				m[i][j] = 0
			}
		}
	}

	return string(s[max_index-max_len : max_index])
}

func longestPalindrome2(s string) string {

	IsParlindrome := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		IsParlindrome[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		IsParlindrome[i][i] = true
	}

	max_i := 0
	max_l := 1
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			if s[j] == s[i+j] {
				if i == 1 {
					IsParlindrome[j][i+j] = true
				} else {
					IsParlindrome[j][i+j] = IsParlindrome[j+1][i+j-1]
				}
				if IsParlindrome[j][i+j] {
					if max_l <= i+1 {
						max_l = i + 1
						max_i = j
					}
				}
			} else {
				IsParlindrome[j][i+j] = false
			}
		}
	}
	return string(s[max_i : max_i+max_l])
}

func main() {
	fmt.Println(longestPalindrome2("aabba"))
}
