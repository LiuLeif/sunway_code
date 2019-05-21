// 2018-10-23 11:13
package main

import "fmt"

func subarrayBitwiseORs(A []int) int {
	result, prev := map[int]bool{}, map[int]bool{}
	for _, n := range A {
		tmp := map[int]bool{}
		tmp[n] = true
		for k, _ := range prev {
			tmp[n|k] = true
		}
		prev = tmp
		for k, _ := range prev {
			result[k] = true
		}
	}
	return len(result)
}

func main() {
	fmt.Println(subarrayBitwiseORs([]int{1, 2, 4}))
}
