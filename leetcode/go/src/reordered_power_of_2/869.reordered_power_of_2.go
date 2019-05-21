// 2018-10-22 15:55
package main

import "fmt"
import "math"

func reorderedPowerOf2(N int) bool {
	powerOfTwos := [][26]int{}
	for i := 1; i < math.MaxInt32; i *= 2 {
		counter := [26]int{}
		x := i
		for x != 0 {
			counter[x%10]++
			x /= 10
		}
		powerOfTwos = append(powerOfTwos, counter)
	}

	target := [26]int{}
	for N != 0 {
		target[N%10]++
		N /= 10
	}
	for _, v := range powerOfTwos {
		if v == target {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(reorderedPowerOf2(46))
}
