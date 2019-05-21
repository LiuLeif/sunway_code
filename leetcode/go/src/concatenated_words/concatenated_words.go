// 2018-02-01 11:56
package main

import (
	"fmt"
)

type Node struct {
	children [26]*Node
	val      string
}

func buildTrie(words []string) *Node {
	root := &Node{}
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			index := int(word[i] - 'a')
			if node.children[index] == nil {
				node.children[index] = &Node{}
			}
			node = node.children[index]
		}
		node.val = word
	}
	return root
}

func test(word string, root *Node, count int) bool {
	node := root
	for i := 0; i < len(word); i++ {
		node = node.children[int(word[i]-'a')]
		if node == nil {
			return false
		}
		if node.val != "" {
			if node.val == word {
				return count > 0
			}
			if test(word[len(node.val):], root, count+1) {
				return true
			}
		}
	}

	return false
}

func findAllConcatenatedWordsInADict(words []string) []string {
	ret := []string{}
	root := buildTrie(words)
	for _, word := range words {
		if test(word, root, 0) {
			ret = append(ret, word)
		}
	}
	return ret
}
func main() {
	fmt.Println(findAllConcatenatedWordsInADict([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}))
}
