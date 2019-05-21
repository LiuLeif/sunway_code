// 2018-11-02 10:11
package main

import "fmt"

func threeEqualParts(A []int) []int {
	// count 1
	ones := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 1 {
			ones++
		}
	}
	if ones%3 != 0 {
		return []int{-1, -1}
	}
	if ones == 0 {
		return []int{0, len(A) - 1}
	}
	step := ones / 3
	index := [3]int{}
	count := 0

	for i := 0; i < len(A); i++ {
		if A[i] == 1 {
			if count%step == 0 {
				index[count/step] = i
			}
			count++
		}
	}
	for index[2] < len(A) {
		if A[index[0]] != A[index[1]] || A[index[0]] != A[index[2]] {
			return []int{-1, -1}
		}
		index[0]++
		index[1]++
		index[2]++
	}
	return []int{index[0] - 1, index[1]}
}

func main() {
	fmt.Println(threeEqualParts([]int{1, 0, 1, 0, 1, 0}))
}
