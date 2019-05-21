// 2017-11-16 11:33
package main

import "fmt"

func dfs(m map[int]bool, stones []int, lastPos, k int) bool {
	if lastPos == stones[len(stones)-1] {
		return true
	}
	for i := 1; i >= -1; i-- {
		if k+i <= 0 {
			continue
		}
		if m[lastPos+k+i] && dfs(m, stones, lastPos+k+i, k+i) {
			return true
		}
	}
	return false
}

func canCross(stones []int) bool {
	for i := 1; i < len(stones); i++ {
		if i > 3 && stones[i]-stones[i-1] > stones[i-1]*2 {
			return false
		}
	}

	m := make(map[int]bool)
	for _, v := range stones {
		m[v] = true
	}

	return dfs(m, stones, 0, 0)
}

func canCrossHash(stones []int) bool {
	m := make(map[int]map[int]bool)
	m[0] = map[int]bool{0: true}

	for i := 0; i < len(stones); i++ {
		x := stones[i]
		for k, _ := range m[x] {
			_, ok := m[x+k]
			if !ok {
				m[x+k] = map[int]bool{}
			}
			m[x+k][k] = true
			_, ok = m[x+k-1]
			if !ok {
				m[x+k-1] = map[int]bool{}
			}
			m[x+k-1][k-1] = true
			_, ok = m[x+k+1]
			if !ok {
				m[x+k+1] = map[int]bool{}
			}
			m[x+k+1][k+1] = true
		}

	}
	return len(m[stones[len(stones)-1]]) != 0
}

func main() {
	fmt.Println(canCrossHash([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	fmt.Println(canCrossHash([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}
