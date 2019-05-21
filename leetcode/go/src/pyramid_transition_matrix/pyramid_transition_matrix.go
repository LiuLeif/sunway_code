// 2018-02-02 16:01
package main

import "fmt"

func getNext(ret *[]string, curr, bottom string, m map[string][]string) {
	if len(bottom) == 1 {
		*ret = append(*ret, curr)
		return
	}
	for _, v := range m[bottom[:2]] {
		getNext(ret, curr+v, bottom[1:], m)
	}
}

func dfs(bottom string, m map[string][]string, visited map[string]bool) bool {
	if len(bottom) == 1 {
		return true
	}
	if visited[bottom] {
		return false
	}
	next := []string{}
	getNext(&next, "", bottom, m)
	for _, v := range next {
		if dfs(v, m, visited) {
			return true
		}
	}
	visited[bottom] = true
	return false
}

func pyramidTransition(bottom string, allowed []string) bool {
	m := map[string][]string{}
	for _, s := range allowed {
		if m[s[:2]] == nil {
			m[s[:2]] = []string{}
		}
		m[s[:2]] = append(m[s[:2]], s[2:])
	}

	visited := map[string]bool{}
	return dfs(bottom, m, visited)
}

func main() {
	fmt.Println(pyramidTransition("XYZ", []string{"XYD", "YZE", "DEA", "FFF"}))
	fmt.Println(pyramidTransition("XXYX", []string{"XXX", "XXY", "XYX", "XYY", "YXZ"}))
}
