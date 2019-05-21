// 2017-11-15 14:25
package main

import "fmt"

func dfs(record []int, ret *[][]int, k, n, start int) {
	if n == 0 {
		if k == 0 {
			*ret = append(*ret, append([]int{}, record...))
		}
		return
	}
	if k == 0 && n != 0 {
		return
	}

	for i := start; i < 10; i++ {
		if n >= i {
			record = append(record, i)
			dfs(record, ret, k-1, n-i, i+1)
			record = record[:len(record)-1]
		}
	}

}

func combinationSum3(k int, n int) [][]int {
	ret := [][]int{}
	dfs([]int{}, &ret, k, n, 1)
	return ret
}
func main() {
	fmt.Println(combinationSum3(2, 10))
}
