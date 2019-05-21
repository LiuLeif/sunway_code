// 2018-01-29 13:17
package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func dupNeeds(needs [6]int) [6]int {
	ret := [6]int{}
	copy(ret[:], needs[:])
	return ret
}

func helper(m map[[6]int]int, price []int, specials [][]int, needs [6]int) int {
	if _, ok := m[needs]; ok {
		return m[needs]
	}
	money := 0
	for i := 0; i < len(price); i++ {
		money += needs[i] * price[i]
	}
	for _, special := range specials {
		tmp := dupNeeds(needs)
		valid := true
		for k := 0; k < len(price); k++ {
			tmp[k] -= special[k]
			if tmp[k] < 0 {
				valid = false
				break
			}
		}
		if valid {
			money = min(money, special[len(special)-1]+helper(m, price, specials, tmp))
		}
	}
	m[needs] = money
	return money
}

func shoppingOffers(price []int, specials [][]int, needs []int) int {
	m := map[[6]int]int{}
	n := [6]int{}
	copy(n[:], needs)
	return helper(m, price, specials, n)
}
func main() {
	fmt.Println(shoppingOffers([]int{2, 5}, [][]int{
		{3, 0, 5},
		{1, 2, 10},
	}, []int{3, 2}))

	fmt.Println(shoppingOffers([]int{2, 3, 4}, [][]int{
		{1, 1, 0, 4},
		{2, 2, 1, 9},
	}, []int{1, 2, 1}))
}
