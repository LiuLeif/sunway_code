// 2018-10-16 12:25
package main

import "fmt"

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			ten++
			if five > 0 {
				five--
			} else {
				return false
			}
		} else {
			if ten > 0 {
				ten--
				if five > 0 {
					five--
				} else {
					return false
				}
			} else {
				if five >= 3 {
					five -= 3
				} else {
					return false
				}
			}
		}
	}
	return true
}
func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
}
