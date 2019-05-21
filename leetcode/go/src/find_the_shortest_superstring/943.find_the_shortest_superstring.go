// 2018-12-06 17:32
package main

import (
	"fmt"
	"strings"
)

func shortestSuperstring(A []string) string {
	N := len(A)
	concat := make([][]string, N)
	for i := 0; i < N; i++ {
		concat[i] = make([]string, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			k := 0
			for ; k < len(A[i]); k++ {
				if strings.HasPrefix(A[j], A[i][k:]) {
					break
				}
			}
			concat[i][j] = A[j][len(A[i])-k:]
		}
	}

	var memo map[[2]int]string

	var dfs func(visited int, last int) string
	dfs = func(visited int, last int) string {
		if cache, ok := memo[[2]int{visited, last}]; ok {
			return cache
		}
		ret := ""
		for i := 0; i < N; i++ {
			if visited&(1<<uint(i)) != 0 {
				continue
			}
			tmp := concat[last][i] + dfs(visited|(1<<uint(i)), i)
			if ret == "" || len(ret) > len(tmp) {
				ret = tmp
			}
		}
		memo[[2]int{visited, last}] = ret
		return ret
	}

	ret := ""
	for i := 0; i < N; i++ {
		memo = map[[2]int]string{}
		str := A[i] + dfs(1<<uint(i), i)
		if ret == "" || len(str) < len(ret) {
			ret = str
		}
	}
	return ret
}

func main() {
	fmt.Println(shortestSuperstring([]string{"alex", "loves", "leetcode"}))
	fmt.Println(shortestSuperstring([]string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"}))
	fmt.Println(shortestSuperstring([]string{"gnwvaraecglrhxsggyai", "raecglrhxsggyaifdn", "bliwndhvprmozdduzzxm", "kvokzckhugsdzbmnf", "ykjkggnwvaraecglr", "sdzbmnfykeywf", "dhvprmozdduzzxms", "ggyaifdnbliwndh", "yaifdnbliwndhvprmoz", "mskvokzckhugsdzbm"}))
}
