// 2018-11-17 13:22
package main

import (
	"fmt"
)

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	ret := [][]int{}
	ret = append(ret, []int{r0, c0})

	count := R*C - 1
	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	d := 0
	prevStep := []int{0, 0}

	for count > 0 {
		step := prevStep[0]
		for i := 0; i < step+1; i++ {
			r0, c0 = r0+directions[d][0], c0+directions[d][1]
			if r0 >= 0 && r0 < R && c0 >= 0 && c0 < C {
				ret = append(ret, []int{r0, c0})
				count--
			}
		}

		d = (d + 1) % 4
		prevStep = append(prevStep[1:], step+1)
	}
	return ret
}

func main() {
	fmt.Println(spiralMatrixIII(5, 6, 1, 4))
	fmt.Println(spiralMatrixIII(1, 4, 0, 0))
}
