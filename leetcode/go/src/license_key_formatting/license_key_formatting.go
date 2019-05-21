// 2018-02-12 13:58
package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
	S = strings.Join(strings.Split(S, "-"), "")
	ret := []string{}
	first := 0
	if len(S)%K != 0 {
		first = len(S) % K
		ret = append(ret, S[:first])
	}
	for i := 0; i < len(S)/K; i++ {
		ret = append(ret, S[first:first+K])
		first += K
	}
	return strings.ToUpper(strings.Join(ret, "-"))
}
func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 2))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
}
