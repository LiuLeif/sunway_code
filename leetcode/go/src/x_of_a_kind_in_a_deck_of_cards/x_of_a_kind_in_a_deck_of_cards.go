// 2018-10-12 22:01
package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {
	if len(deck) == 0 {
		return false
	}
	counter := map[int]int{}
	for _, v := range deck {
		counter[v]++
	}

	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	gcdValue := 0
	for _, v := range counter {
		if gcdValue == 0 {
			gcdValue = v
		} else {
			gcdValue = gcd(gcdValue, v)
		}
		if gcdValue == 1 {
			break
		}
	}
	if gcdValue == 1 {
		return false
	} else {
		return true
	}
}
func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
}
