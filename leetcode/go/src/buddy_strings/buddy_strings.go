// 2018-10-12 15:45
package main

import "fmt"

func buddyStrings(A string, B string) bool {
	count := 0
	a, b := byte(0), byte(0)
	for i := 0; i < len(A); i++ {
		if A[i] == B[i] {
			continue
		}
		if count == 0 {
			a, b = A[i], B[i]
		} else if count == 1 {
			if a != B[i] || b != A[i] {
				return false
			}
		}
		count++
	}

	isDuplicate := func() bool {
		mark := [26]bool{}
		for _, v := range A {
			if mark[v-97] {
				return true
			}
			mark[v-97] = true
		}
		return false
	}

	if count == 0 {
		return isDuplicate()
	} else if count == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb"))
	fmt.Println(buddyStrings("aaxc", "axab"))
}
