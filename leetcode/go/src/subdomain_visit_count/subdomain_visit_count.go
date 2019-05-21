// 2018-04-08 16:23
package main

import (
	"fmt"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := map[string]int{}
	for _, cpdomain := range cpdomains {
		count := 0
		domain := ""
		fmt.Sscanf(cpdomain, "%d %s", &count, &domain)
		dot := -1
		for {
			domain = domain[dot+1:]
			m[domain] += count
			dot = strings.Index(domain, ".")
			if dot == -1 {
				break
			}
		}
	}
	ret := []string{}
	for domain, count := range m {
		ret = append(ret, fmt.Sprintf("%d %s", count, domain))
	}
	return ret
}
func main() {
	fmt.Println(subdomainVisits([]string{
		"900 google.mail.com",
		"50 yahoo.com",
		"1 intel.mail.com",
		"5 wiki.org",
	}))
}
