// 2018-02-06 12:01
package main

import "fmt"

func helper(mem map[[2]int]bool, s, p string, x, y int) bool {
	if len(s) == x && len(p) == y {
		return true
	}
	if v, ok := mem[[2]int{x, y}]; ok {
		return v

	}
	if y < len(p)-1 && p[y+1] == '*' && helper(mem, s, p, x, y+2) {
		mem[[2]int{x, y}] = true
		return true
	}
	if len(s) == x || len(p) == y {
		mem[[2]int{x, y}] = false
		return false
	}
	if p[y] == '*' {
		mem[[2]int{x, y}] = false
		return false
	}

	if s[x] == p[y] || p[y] == '.' {
		if helper(mem, s, p, x+1, y+1) {
			mem[[2]int{x, y}] = true
			return true
		}
		if y < len(p)-1 && p[y+1] == '*' {
			if helper(mem, s, p, x+1, y) {
				mem[[2]int{x, y}] = true
				return true
			}
		}
	}
	mem[[2]int{x, y}] = false
	return false
}

func isMatch(s string, p string) bool {
	mem := map[[2]int]bool{}
	return helper(mem, s, p, 0, 0)
}
func main() {
	fmt.Println(isMatch("", ".*"))
}
