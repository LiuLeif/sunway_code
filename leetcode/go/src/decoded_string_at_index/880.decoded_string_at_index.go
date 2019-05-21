// 2018-10-26 09:39
package main

import "fmt"

func decodeAtIndex(S string, K int) string {
	size := 0
	for i := 0; i < len(S); i++ {
		if S[i] >= 48 && S[i] <= 57 {
			// digit
			size *= int(S[i] - 48)
		} else {
			size++
		}
	}
	for i := len(S) - 1; i >= 0; i-- {
		K %= size
		if K == 0 && !(S[i] >= 48 && S[i] <= 57) {
			return string(S[i])
		}

		if S[i] >= 48 && S[i] <= 57 {
			// digit
			size /= int(S[i] - 48)
		} else {
			size--
		}
	}
	return ""
}

func main() {
	fmt.Println(decodeAtIndex("a23", 5))
}
