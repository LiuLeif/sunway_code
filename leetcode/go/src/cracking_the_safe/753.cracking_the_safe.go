// 2018-11-22 15:38
package main

import (
	"fmt"
)

func Pow(x, y int) int {
	ret := 1
	for i := 0; i < x; i++ {
		ret *= y
	}
	return ret
}

func crackSafe(n int, k int) string {
	visited := map[string]bool{}
	N := Pow(n, k)
	var traverse func(prev string) string
	traverse = func(prev string) string {
		if len(visited) == N {
			return prev
		}
		for i := 0; i < k; i++ {
			c := string(i + '0')
			n := prev[len(prev)-n+1:] + c
			if visited[n] {
				continue
			}
			visited[n] = true
			tmp := traverse(prev + c)
			if len(tmp) != 0 {
				return tmp
			}
			delete(visited, n)
		}
		return ""
	}

	ret := ""
	for i := 0; i < n-1; i++ {
		ret += "0"
	}
	ret = traverse(ret)
	return ret
}

func main() {
	fmt.Println(crackSafe(2, 4)) // 00102031121322330
}
