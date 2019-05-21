// 2017-12-28 14:52
package main

import "fmt"

func dumpResult(record []string, top, begin string, parent map[string]map[string]bool, ret *[][]string, minLen int) {

	if top == begin {
		if len(record) == minLen {
			record = append([]string{begin}, record...)
			*ret = append(*ret, append([]string{}, record...))
		}
		return
	}
	record = append([]string{top}, record...)
	for v, _ := range parent[top] {
		dumpResult(record, v, begin, parent, ret, minLen)
	}
	record = record[:len(record)-1]
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ret := [][]string{}
	queue := []string{beginWord, ""}

	words := map[string]bool{}
	for i := 0; i < len(wordList); i++ {
		words[wordList[i]] = true
	}
	parent := map[string]map[string]bool{}
	visited := map[string]bool{}
	levelVisited := map[string]bool{}

	minLen, count := 0, 0
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == endWord {
			minLen = count
			dumpResult([]string{}, endWord, beginWord, parent, &ret, minLen)
			break
		}
		if top == "" {
			queue = append(queue, "")
			count++
			levelVisited = map[string]bool{}
			continue
		}
		for i := 0; i < len(top); i++ {
			for _, v := range "abcdefghijklmnopqrstuvwxyz" {
				t := top[:i] + string(v) + top[i+1:]
				if words[t] && !visited[t] {
					if parent[t] == nil {
						parent[t] = map[string]bool{}
					}
					if t != top {
						parent[t][top] = true
					}

					queue = append(queue, t)
					levelVisited[t] = true
					if t != endWord {
						visited[t] = true
					}
				} else if levelVisited[t] {
					if parent[t] == nil {
						parent[t] = map[string]bool{}
					}
					if t != top {
						parent[t][top] = true
					}
				}
			}
		}
	}
	return ret
}
func main() {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(findLadders("hit", "log", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(findLadders("a", "c", []string{"a", "b", "c"}))
	fmt.Println(findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))
}
