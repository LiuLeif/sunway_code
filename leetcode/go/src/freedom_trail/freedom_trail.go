// 2018-02-27 10:04
package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func distance(f, t, l int) int {
	if f == t {
		return 0
	}
	if f > t {
		f, t = t, f
	}
	ret := min(t-f, f+l-t)
	return ret
}

func dfs(mem map[pair]int, m map[byte][]int, ring, key string, start int) int {
	if len(key) == 0 {
		return 0
	}
	if cache, ok := mem[pair{start, key}]; ok {
		return cache
	}
	ret := math.MaxInt32
	for _, pos := range m[key[0]] {
		ret = min(ret, distance(start, pos, len(ring))+dfs(mem, m, ring, key[1:], pos))
	}
	mem[pair{start, key}] = ret
	return ret
}

type pair struct {
	start int
	key   string
}

func findRotateSteps(ring string, key string) int {
	distances := map[byte][]int{}
	for i := 0; i < len(ring); i++ {
		if distances[ring[i]] == nil {
			distances[ring[i]] = []int{}
		}
		distances[ring[i]] = append(distances[ring[i]], i)
	}
	mem := map[pair]int{}
	return dfs(mem, distances, ring, key, 0) + len(key)
}
func main() {
	fmt.Println(findRotateSteps("godding", "gd"))
}
