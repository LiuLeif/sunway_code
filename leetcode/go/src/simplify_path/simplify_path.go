// 2018-01-19 14:54
package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	stack := []string{}
	for _, p := range paths {
		switch p {
		case ".", "":
			continue
		case "..":
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, p)
		}
	}

	ret := strings.Join(stack, "/")
	ret = strings.TrimRight(ret, "/")
	return "/" + ret
}
func main() {
	fmt.Println(simplifyPath("/home/foo/..//../.."))
	fmt.Println(simplifyPath("/a/./b///../c/../././../d/..//../e/./f/./g/././//.//h///././/..///"))
}
