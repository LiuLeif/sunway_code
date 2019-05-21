// 2018-01-02 13:51
package main

import "fmt"

func helper(mem map[Pair]int, remain int, coins []int, start int) int {
	if remain == 0 {
		return 1
	}
	if t, ok := mem[Pair{remain, start}]; ok {
		return t
	}
	ret := 0
	for i := start; i < len(coins); i++ {
		if remain-coins[i] >= 0 {
			ret += helper(mem, remain-coins[i], coins, i)
		}
	}
	mem[Pair{remain, start}] = ret
	return ret
}

type Pair struct {
	remain, start int
}

func change(amount int, coins []int) int {
	mem := make(map[Pair]int)
	return helper(mem, amount, coins, 0)
}

func main() {
	fmt.Println(change(500, []int{3, 5, 7, 8, 9, 10, 11}))
}
