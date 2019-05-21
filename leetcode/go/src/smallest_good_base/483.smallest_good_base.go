// 2018-11-13 09:35
package main

import (
	"fmt"
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	// 13=1(1-3^3)/1-3=13=(q^n-1)/(q-1)
	N, _ := strconv.ParseUint(n, 10, 64)
	check := func(k int) uint64 {
		var lo uint64 = 2
		var hi uint64
		if k == 1 {
			hi = N
		} else {
			hi = uint64(math.Pow(float64(N), 1.0/float64(k))) + 1
		}

		for lo <= hi {
			mid := (lo + hi) / 2
			sum := uint64(0)
			curr := uint64(1)
			for i := 0; i <= k; i++ {
				sum += curr
				curr *= mid
			}
			if N == sum {
				return mid
			} else if sum < N {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return 0
	}

	for k := 62; k >= 1; k-- {
		ret := check(k)
		if ret != 0 {
			t := strconv.FormatUint(ret, 10)
			return t
		}
	}
	return "no"
}

func main() {
	fmt.Println(smallestGoodBase("13"))
	fmt.Println(smallestGoodBase("4681"))
	fmt.Println(smallestGoodBase("919818571896748567"))
	fmt.Println(smallestGoodBase("2251799813685247"))
}
