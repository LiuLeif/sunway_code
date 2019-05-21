// 2018-10-30 11:14
package main

import "fmt"

func numUniqueEmails(emails []string) int {
	m := map[string]bool{}
	for _, s := range emails {
		s2 := ""
		skip := false
		for i := 0; i < len(s); i++ {
			if s[i] == '.' {
				continue
			}
			if s[i] == '+' {
				skip = true
				continue
			}
			if s[i] == '@' {
				s2 += s[i:]
				break
			}
			if !skip {
				s2 += string(s[i])
			}
		}
		m[s2] = true
	}
	return len(m)
}

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
