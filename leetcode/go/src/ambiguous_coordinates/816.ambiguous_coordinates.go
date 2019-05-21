// 2018-11-08 11:30
package main

import "fmt"

func ambiguousCoordinates(S string) []string {
	S = S[1 : len(S)-1]
	ret := []string{}
	parse := func(s string) []string {
		// 12, 000, 001, 010
		ret := []string{}
		if s[0] != '0' || (s[0] == '0' && len(s) == 1) {
			ret = append(ret, s)
		}
		if s[len(s)-1] != '0' {
			// a.b
			if s[0] == '0' {
				ret = append(ret, s[:1]+"."+s[1:])
			} else {
				for i := 1; i < len(s); i++ {
					ret = append(ret, s[:i]+"."+s[i:])
				}
			}
		}
		return ret
	}

	for i := 1; i < len(S); i++ {
		for _, x := range parse(S[:i]) {
			for _, y := range parse(S[i:]) {
				ret = append(ret, fmt.Sprintf("(%s, %s)", x, y))
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(ambiguousCoordinates("(100)"))
}
