// 2017-11-24 17:37
package main

import "fmt"

func dfs(ret *[][]int, record []int, n, k, start int) {
	if len(record) == k {
		*ret = append(*ret, append([]int{}, record...))
		return
	}
	for i := start; i < n+1; i++ {
		record = append(record, i)
		dfs(ret, record, n, k, i+1)
		record = record[:len(record)-1]
	}
}

func combine(n int, k int) [][]int {
	ret := [][]int{}
	dfs(&ret, []int{}, n, k, 1)
	return ret
}
func main() {
	fmt.Println(combine(4, 3))
}
