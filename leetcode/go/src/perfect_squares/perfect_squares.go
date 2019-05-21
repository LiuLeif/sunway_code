// 2017-12-19 15:43
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

func helper(m map[int]int, n int) int {
	if n <= 1 {
		return n
	}
	if m[n] != 0 {
		return m[n]
	}
	i := 1
	ret := math.MaxInt32
	for i*i <= n {
		ret = min(ret, helper(m, n-i*i)+1)
		i += 1
	}
	m[n] = ret
	return ret
}

func numSquares(n int) int {
	return helper(map[int]int{}, n)
}

func numSquaresBFS(n int) int {
	// LTE for numSquares(7168)
	queue := []int{n, -1}
	count := 0
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == -1 {
			count++
			queue = append(queue, -1)
			continue
		}
		if top == 0 {
			return count
		}
		i := 1
		for i*i <= top {
			queue = append(queue, top-i*i)
			i += 1
		}
	}
	return 0
}

func main() {
	fmt.Println(numSquares(7168))
}
