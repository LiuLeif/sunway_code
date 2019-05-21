// 2018-01-10 10:45
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func helper(ret *[]string, record []string, s string) {
	if len(s) == 0 && len(record) == 4 {
		*ret = append(*ret, strings.Join(record, "."))
		return
	}
	if len(record) == 4 || len(s) == 0 {
		return
	}
	if s[0] == '0' {
		record = append(record, "0")
		helper(ret, record, s[1:])
		record = record[:len(record)-1]
	} else {
		for i := 0; i < min(3, len(s)); i++ {
			x := s[:i+1]
			ival, _ := strconv.Atoi(x)
			if ival <= 255 {
				record = append(record, x)
				helper(ret, record, s[i+1:])
				record = record[:len(record)-1]
			}
		}
	}
}

func restoreIpAddresses(s string) []string {
	ret := []string{}
	helper(&ret, []string{}, s)
	return ret
}
func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("1141001"))
}
