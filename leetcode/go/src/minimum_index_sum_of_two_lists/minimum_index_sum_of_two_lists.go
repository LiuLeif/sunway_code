// 2017-12-29 14:47
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

func findRestaurant(list1 []string, list2 []string) []string {
	m, sum := map[string]int{}, map[string]int{}
	for k, v := range list1 {
		m[v] = k
	}

	minimum := math.MaxInt32
	for k, v := range list2 {
		if x, ok := m[v]; ok {
			sum[v] = x + k
			minimum = min(minimum, x+k)
		}
	}

	ret := []string{}
	for k, v := range sum {
		if v == minimum {
			ret = append(ret, k)
		}
	}

	return ret
}
func main() {
	fmt.Println(findRestaurant(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"},
	))
}
