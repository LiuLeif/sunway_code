// 2017-12-08 16:55
package main

import "fmt"

func dfs(ret *[]string, s string, left, right int) {
	if left == 0 && right == 0 {
		*ret = append(*ret, s)
		return
	}
	if left > 0 {
		dfs(ret, s+"(", left-1, right)
	}
	if right > left {
		dfs(ret, s+")", left, right-1)
	}
}

func generateParenthesis(n int) []string {
	// ()()()
	ret := []string{}
	dfs(&ret, "", n, n)
	return ret
}
func main() {
	fmt.Println(generateParenthesis(3))
}
