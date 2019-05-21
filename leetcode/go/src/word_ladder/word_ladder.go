// 2017-11-14 18:54
package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue := make([]string, 2)
	queue[0] = beginWord
	queue[1] = ""
	words := make(map[string]bool)
	for _, v := range wordList {
		words[v] = true
	}

	level := 1
	for len(queue) != 1 {
		k := queue[0]
		if k == "" {
			level++
			queue = append(queue, "")
		}
		if k == endWord {
			return level
		}
		queue = queue[1:]
		for i := 0; i < len(k); i++ {
			for _, j := range "abcdefghijklmnopqrstuvwxyz" {
				k2 := k[:i] + string(j) + k[i+1:]
				if words[k2] == true {
					words[k2] = false
					queue = append(queue, k2)
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("abc", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
