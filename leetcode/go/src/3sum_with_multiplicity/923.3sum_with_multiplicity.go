// 2018-10-23 15:25
package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumMulti(A []int, target int) int {
	MOD := int(math.Pow(10, 9)) + 7
	sort.Ints(A)
	counter := map[int]int{}
	for _, v := range A {
		counter[v]++
	}
	compute := func(a, b, c int) int {
		ca, cb, cc := counter[a], counter[b], counter[c]
		if a == b && b == c {
			return ca * (ca - 1) * (ca - 2) / 6
		} else if a == b {
			return ca * (ca - 1) * cc / 2
		} else if b == c {
			return cb * (cb - 1) * ca / 2
		} else {
			return ca * cb * cc
		}
		return 0
	}

	ret := 0
	for i := 0; i < len(A)-2; i++ {
		if i != 0 && A[i] == A[i-1] {
			continue
		}
		lo, hi := i+1, len(A)-1
		for lo < hi {
			if A[lo]+A[hi] == target-A[i] {
				ret += compute(A[i], A[lo], A[hi])
				ret %= MOD
				for lo < hi && A[lo] == A[lo+1] {
					lo++
				}
				for lo < hi && A[hi] == A[hi-1] {
					hi--
				}
				lo++
				hi--
			} else if A[lo]+A[hi] > target-A[i] {
				hi--
			} else {
				lo++
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(threeSumMulti([]int{0, 0, 0, 0}, 0))
}
