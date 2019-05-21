// 2018-01-15 14:35
package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	children [26]*TrieNode
	word     string
}

func buildTrie(dict []string) *TrieNode {
	root := &TrieNode{}
	for i := 0; i < len(dict); i++ {
		word := dict[i]
		curr := root
		for j := 0; j < len(word); j++ {
			if curr.children[word[j]-'a'] == nil {
				curr.children[word[j]-'a'] = &TrieNode{}
			}
			curr = curr.children[word[j]-'a']
		}
		curr.word = word
	}
	return root
}

func searchTrie(root *TrieNode, word string) string {
	curr := root
	for i := 0; i < len(word); i++ {
		if curr.word != "" {
			return curr.word
		}
		if curr.children[word[i]-'a'] == nil {
			return word
		}
		curr = curr.children[word[i]-'a']
	}

	return word
}

func replaceWords(dict []string, sentence string) string {
	root := buildTrie(dict)
	tokens := strings.Split(sentence, " ")
	for i := 0; i < len(tokens); i++ {
		tokens[i] = searchTrie(root, tokens[i])
	}
	return strings.Join(tokens, " ")
}
func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}
