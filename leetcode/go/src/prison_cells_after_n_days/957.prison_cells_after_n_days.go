// 2018-12-18 16:09
package main

import "fmt"

func prisonAfterNDays(cells []int, N int) []int {
	cellsArray := [8]int{}
	copy(cellsArray[:], cells)

	next := func() [8]int {
		ret := [8]int{}
		copy(ret[:], cellsArray[:])
		for i := 0; i < 8; i++ {
			if i == 0 || i == 7 {
				ret[i] = 0
			} else {
				if cellsArray[i-1] == cellsArray[i+1] {
					ret[i] = 1
				} else {
					ret[i] = 0
				}
			}
		}
		return ret
	}

	cycle := -1
	cycleStart := 0
	m := map[[8]int]int{}
	for i := 1; i < N+1; i++ {
		cellsArray = next()
		if prev, ok := m[cellsArray]; ok {
			cycle = i - prev
			cycleStart = prev
			break
		}
		m[cellsArray] = i
	}

	if cycle == -1 {
		return cellsArray[:]
	}

	N = (N-cycleStart)%cycle + cycleStart

	for k, v := range m {
		if v == N {
			return k[:]
		}
	}
	return []int{}
}

func main() {
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 7))
	fmt.Println(prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000))
	fmt.Println(prisonAfterNDays([]int{1, 1, 0, 1, 1, 0, 1, 1}, 6))
}
