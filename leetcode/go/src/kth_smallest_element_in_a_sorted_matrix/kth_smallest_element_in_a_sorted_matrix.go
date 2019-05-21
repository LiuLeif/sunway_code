// 2017-11-14 18:54
package main

import "fmt"

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left := matrix[0][0]
	right := matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		total := 0
		for i := 0; i < n; i++ {
			if matrix[i][0] > mid {
				break
			}
			for j := 0; j < n; j++ {
				if matrix[i][j] <= mid {
					total++
				} else {
					break
				}
			}
		}
		if total >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func main() {
	matrix := make([][]int, 3)
	matrix[0] = []int{1, 5, 12}
	matrix[1] = []int{10, 11, 13}
	matrix[2] = []int{12, 13, 15}
	fmt.Println(kthSmallest(matrix, 8))
}
