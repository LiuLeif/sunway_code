// 2018-10-23 14:11
package main

import (
	"fmt"
	"math"
	"strconv"
)

func backtrack(S string, record []int) []int {
	if len(S) == 0 {
		if len(record) >= 3 {
			return record
		} else {
			return nil
		}
	}
	for i := 1; i < len(S)+1; i++ {
		if S[0] == '0' && i != 1 {
			break
		}
		tmp, err := strconv.Atoi(S[:i])
		if err != nil || tmp > math.MaxInt32 {
			break
		}
		if len(record) >= 2 && tmp != record[len(record)-1]+record[len(record)-2] {
			continue
		}
		ret := backtrack(S[i:], append(record, tmp))
		if ret != nil {
			return ret
		}
	}
	return nil
}

func splitIntoFibonacci(S string) []int {
	return backtrack(S, []int{})
}

func main() {
	fmt.Println(splitIntoFibonacci("123456579"))
}
