// 2018-11-06 10:16
package main

import "fmt"

type Tuple struct {
	a, b, c int
}

func shortestBridge(A [][]int) int {
	N := len(A)
	queue := []Tuple{}
	queue2 := []Tuple{}
loop:
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if A[i][j] == 1 {
				queue = append(queue, Tuple{i, j, 0})
				queue2 = append(queue, Tuple{i, j, 0})
				A[i][j] = 2
				break loop
			}
		}
	}
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		for _, x := range []int{-1, 0, 1} {
			for _, y := range []int{-1, 0, 1} {
				if (x == 0 && y == 0) || x*y != 0 {
					continue
				}
				a, b := top.a+x, top.b+y
				if a >= 0 && a < N && b >= 0 && b < N && A[a][b] == 1 {
					fmt.Println(a, b)
					A[a][b] = 2
					queue = append(queue, Tuple{a, b, 0})
					queue2 = append(queue2, Tuple{a, b, 0})
				}
			}
		}
	}

	for len(queue2) != 0 {
		top := queue2[0]
		queue2 = queue2[1:]
		for _, x := range []int{-1, 0, 1} {
			for _, y := range []int{-1, 0, 1} {
				if (x == 0 && y == 0) || x*y != 0 {
					continue
				}
				a, b := top.a+x, top.b+y
				if a >= 0 && a < N && b >= 0 && b < N {
					if A[a][b] == 1 {
						return top.c
					}
					if A[a][b] == 0 {
						A[a][b] = 2
						queue2 = append(queue2, Tuple{a, b, top.c + 1})
					}
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(shortestBridge([][]int{
		{0, 1},
		{1, 0},
	}))
}
