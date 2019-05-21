// 2018-12-05 15:22
package main

import "fmt"

func numSimilarGroups(A []string) int {
	set := map[string]bool{}
	for _, s := range A {
		set[s] = true
	}

	A = []string{}
	for s := range set {
		A = append(A, s)
	}

	N := len(A)
	parent := make([]int, N)
	depth := make([]int, N)
	for i := 0; i < N; i++ {
		parent[i] = i
	}

	isSimilar := func(a, b string) bool {
		diff := 0
		counter := map[byte]bool{}
		for i := 0; i < len(a); i++ {
			counter[a[i]] = true
			if a[i] != b[i] {
				diff++
			}
			if diff > 2 {
				return false
			}
		}
		return diff == 2 || (diff == 0 && len(counter) != len(a))
	}

	find := func(a int) int {
		for parent[a] != a {
			a = parent[a]
		}
		return a
	}

	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa == pb {
			return
		}
		if depth[pa] > depth[pb] {
			parent[pb] = pa
		} else if depth[pa] < depth[pb] {
			parent[pa] = pb
		} else {
			parent[pa] = pb
			depth[pb]++
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if isSimilar(A[i], A[j]) {
				union(i, j)
			}
		}
	}

	counter := map[int]bool{}
	for i := 0; i < N; i++ {
		counter[find(i)] = true
	}

	return len(counter)
}

func main() {
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"}))
	fmt.Println(numSimilarGroups([]string{"aaa", "aaa", "aaa", "aaa"}))
}
