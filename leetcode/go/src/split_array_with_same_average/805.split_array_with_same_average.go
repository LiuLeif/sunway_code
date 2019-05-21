// 2018-11-06 11:22
package main

import "fmt"

type Tuple struct {
	a, b, c int
}

func splitArraySameAverage(A []int) bool {
	sum := 0
	N := 0
	for _, v := range A {
		sum += v
		N++
	}
	cache := map[Tuple]bool{}

	var dfs func(target, k, index int) bool
	dfs = func(target, k, index int) bool {
		if target == 0 {
			return k == 0
		}
		if cache[Tuple{target, k, index}] {
			return false
		}
		for i := index; i < len(A); i++ {
			if target >= A[i] && dfs(target-A[i], k-1, i+1) {
				return true
			}
		}
		cache[Tuple{target, k, index}] = true
		return false
	}
	for i := 1; i < len(A)/2+1; i++ {
		if sum*i%N != 0 {
			continue
		}
		if dfs(sum*i/N, i, 0) {
			return true
		}
	}
	return false
}

func main() {
	// 1,4,5,8
	// 2,3,6,7
	fmt.Println(splitArraySameAverage([]int{
		60, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	}))
}
