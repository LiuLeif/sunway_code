// 2018-01-27 16:24
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func countOfAtoms(formula string) string {
	stack := []map[string]int{}
	m := map[string]int{}
	for i := 0; i < len(formula); {
		switch formula[i] {
		case '(':
			stack = append(stack, m)
			m = map[string]int{}
			i++
		case ')':
			num := 0
			i++
			for i < len(formula) && formula[i] >= '0' && formula[i] <= '9' {
				num = num*10 + int(formula[i]-'0')
				i++
			}
			if num == 0 {
				num = 1
			}
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for k, v := range m {
				prev[k] += v * num
			}
			m = prev
		default:
			start := i
			i++

			for i < len(formula) && formula[i] >= 'a' && formula[i] <= 'z' {
				i++
			}
			s := formula[start:i]
			num := 0
			for i < len(formula) && formula[i] >= '0' && formula[i] <= '9' {
				num = num*10 + int(formula[i]-'0')
				i++
			}
			if num == 0 {
				num = 1
			}
			m[s] += num
		}
	}

	items := []string{}
	for k, _ := range m {
		items = append(items, k)
	}
	sort.Strings(items)
	ret := ""
	for _, v := range items {
		ret += v
		if m[v] != 1 {
			ret += strconv.Itoa(m[v])
		}
	}

	return ret
}
func main() {
	fmt.Println(countOfAtoms("A2CB2"))
	fmt.Println(countOfAtoms("((N42)24(OB40Li30CHe3O48LiNN26)33(C12Li48N30H13HBe31)21(BHN30Li26BCBe47N40)15(H5)16)14"))
}
