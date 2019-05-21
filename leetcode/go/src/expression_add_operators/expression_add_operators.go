// 2017-12-12 16:01
package main

import (
	"fmt"
	"strconv"
)

func dfs(ret *[]string, remain string, target int, path string, eval, mul int) {
	if len(remain) == 0 && eval == target {
		*ret = append(*ret, path)
		return
	}
	if len(remain) == 0 {
		return
	}
	for i := 1; i < len(remain)+1; i++ {
		x := remain[:i]
		if len(x) > 1 && x[0] == '0' {
			continue
		}
		curr, _ := strconv.Atoi(remain[:i])
		if len(path) == 0 {
			dfs(ret, remain[i:], target, remain[:i], curr, curr)
		} else {
			dfs(ret, remain[i:], target, path+"+"+remain[:i], eval+curr, curr)
			dfs(ret, remain[i:], target, path+"-"+remain[:i], eval-curr, -curr)
			dfs(ret, remain[i:], target, path+"*"+remain[:i], eval-mul+mul*curr, mul*curr)
		}
	}
}
func addOperators(num string, target int) []string {
	ret := []string{}
	dfs(&ret, num, target, "", 0, 0)
	return ret
}
func main() {
	fmt.Println(addOperators("105", 5))
}
