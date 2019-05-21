// 2018-01-02 13:51
package main

import (
	"fmt"
	"math"
	"sort"
)

func primeFactors(n int) []int {
	ret := []int{}
	sqrt := int(math.Sqrt(float64(n)))
	count := make(map[int]int)

	for i := 2; i <= sqrt; i++ {
		for n%i == 0 {
			count[i] += 1
			n = n / i
		}
	}
	if n != 1 {
		count[n] += 1
	}
	fmt.Println(count)
	for k, _ := range count {
		ret = append(ret, k)
	}
	sort.Ints(ret)
	return ret
}

func main() {
	fmt.Println(primeFactors(129))
}
