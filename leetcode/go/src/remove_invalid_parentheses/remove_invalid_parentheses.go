// 2017-11-29 13:49
package main

import "fmt"

func isValid(s string) bool {
	stack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, 0)
		} else if s[i] == ')' {
			if len(stack) != 0 && stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func bfs(ret *[]string, record string) {
	queue := []string{record, "null"}
	found := false
	visited := make(map[string]bool)
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == "null" {
			if found {
				return
			}
			queue = append(queue, "null")
			continue
		}
		if visited[top] {
			continue
		}
		visited[top] = true
		if isValid(top) {
			*ret = append(*ret, top)
			found = true
			continue
		}
		for i := 0; i < len(top); i++ {
			t := top[:i] + top[i+1:]
			queue = append(queue, t)
		}
	}
}

func removeInvalidParentheses(s string) []string {
	ret := []string{}
	bfs(&ret, s)
	return ret
}
func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
}
