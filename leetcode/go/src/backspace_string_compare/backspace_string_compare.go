// 2018-10-15 14:17
package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	stack_s := []rune{}
	stack_t := []rune{}
	for _, v := range S {
		if v == '#' {
			if len(stack_s) != 0 {
				stack_s = stack_s[:len(stack_s)-1]
			}
		} else {
			stack_s = append(stack_s, v)
		}
	}
	for _, v := range T {
		if v == '#' {
			if len(stack_t) != 0 {
				stack_t = stack_t[:len(stack_t)-1]
			}
		} else {
			stack_t = append(stack_t, v)
		}
	}
	return string(stack_s) == string(stack_t)
}

func backspaceCompare2(S string, T string) bool {
	make_iterator := func(s string) func() byte {
		index := len(s)
		skip := 0
		return func() byte {
			for index--; index >= 0; index-- {
				if s[index] == '#' {
					skip++
				} else {
					if skip > 0 {
						skip--
					} else {
						return s[index]
					}
				}
			}
			return 0
		}
	}

	iter_s := make_iterator(S)
	iter_t := make_iterator(T)

	for {
		x := iter_s()
		y := iter_t()

		if x != y {
			return false
		}
		if x == 0 {
			return true
		}
	}
}

func main() {
	fmt.Println(backspaceCompare2("abc", "ad#c"))
}
