// 2017-12-25 17:35
package main

import (
	"fmt"
	"math"
	"sort"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

type Pair struct {
	x, y int
}

func helper(visited map[Pair]int, pairs [][]int, pair Pair) int {
	if visited[pair] != 0 {
		return visited[pair]
	}
	ret := 1
	for i := 0; i < len(pairs); i++ {
		t := Pair{pairs[i][0], pairs[i][1]}
		if t.x > pair.y {
			ret = max(ret, 1+helper(visited, pairs, t))
		}
	}
	visited[pair] = ret
	return ret
}

func findLongestChain(pairs [][]int) int {
	visited := make(map[Pair]int)
	ret := 0
	for _, v := range pairs {
		pair := Pair{v[0], v[1]}
		if visited[pair] == 0 {
			ret = max(ret, helper(visited, pairs, pair))
		}
	}
	return ret
}

func findLongestChainGreedy(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	curr, ret := math.MinInt32, 0
	for i := 0; i < len(pairs); i++ {
		if curr < pairs[i][0] {
			curr, ret = pairs[i][1], ret+1
		}
	}
	return ret
}

func main() {
	pairs := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(findLongestChainGreedy(pairs))
}
