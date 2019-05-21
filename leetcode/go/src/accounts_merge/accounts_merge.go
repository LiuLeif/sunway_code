// 2017-12-27 17:05
package main

import (
	"fmt"
	"sort"
)

type UFS struct {
	parent map[string]string
	depth  map[string]int
}

func MakeUFS() *UFS {
	ufs := &UFS{map[string]string{}, map[string]int{}}
	return ufs
}

func (this *UFS) union(s1, s2 string) {
	s1, s2 = this.find(s1), this.find(s2)
	if this.depth[s1] > this.depth[s2] {
		this.parent[s2] = s1
	} else if this.depth[s1] < this.depth[s2] {
		this.parent[s1] = s2
	} else {
		this.parent[s1] = s2
		this.depth[s2]++
	}
}

func (this *UFS) find(s string) string {
	if this.parent[s] == "" {
		this.parent[s] = s
	}
	for s != this.parent[s] {
		s = this.parent[s]
	}
	return s
}

func accountsMerge(accounts [][]string) [][]string {
	ufs := MakeUFS()
	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			ufs.union(accounts[i][j], accounts[i][1])
		}
	}
	m := map[string][]string{}
	visited := map[string]bool{}
	names := map[string]string{}
	for i := 0; i < len(accounts); i++ {
		name := accounts[i][0]
		for j := 1; j < len(accounts[i]); j++ {
			if visited[accounts[i][j]] {
				continue
			}
			visited[accounts[i][j]] = true
			root := ufs.find(accounts[i][j])
			names[root] = name
			if m[root] == nil {
				m[root] = []string{}
			}
			m[root] = append(m[root], accounts[i][j])
		}
	}
	ret := [][]string{}
	for k, v := range m {
		sort.Strings(v)
		ret = append(ret, append([]string{names[k]}, v...))
	}
	return ret
}

func main() {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	fmt.Println(accountsMerge(accounts))
}
