// 2017-12-12 16:01
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	numStrs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		numStrs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(numStrs, func(i, j int) bool {
		a, b := numStrs[i], numStrs[j]
		x := a + b
		y := b + a
		return x >= y
	})
	ret := strings.Join(numStrs, "")
	ret = strings.TrimLeft(ret, "0")
	if len(ret) == 0 {
		return "0"
	} else {
		return ret
	}
}
func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
