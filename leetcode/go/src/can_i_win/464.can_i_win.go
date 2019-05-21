// 2018-10-29 14:00
package main

import "fmt"

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	// 1 2 3 4 5 6 7 8 9 10
	if desiredTotal == 0 {
		return true
	}
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	var helper func(mask, maxInt, totol int) bool
	memo := map[[2]int]bool{}

	helper = func(mask int, maxInt, total int) bool {
		if total <= 0 {
			return false
		}
		if cache, ok := memo[[2]int{mask, total}]; ok {
			return cache
		}
		for i := 1; i <= maxInt; i++ {
			if mask&(1<<uint(i)) != 0 {
				continue
			}
			if !helper(mask|(1<<uint(i)), maxInt, total-i) {
				memo[[2]int{mask, total}] = true
				return true
			}

		}
		memo[[2]int{mask, total}] = false
		return false
	}
	return helper(0, maxChoosableInteger, desiredTotal)
}

func main() {
	fmt.Println(canIWin(5, 50))
}
