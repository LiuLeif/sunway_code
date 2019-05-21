// 2018-04-08 11:55
package main

import (
	"fmt"
)

type Trie struct {
	children [26]*Trie
	leaf     []int
}

func NewTrieNode() *Trie {
	ret := Trie{}
	ret.leaf = []int{}
	return &ret
}

func BuildTrie(words []string, isForward bool) *Trie {
	root := NewTrieNode()
	for k, word := range words {
		curr := root
		for i := 0; i < len(word); i++ {
			curr.leaf = append(curr.leaf, k)
			c := byte('a')
			if isForward {
				c = word[i]
			} else {
				c = word[len(word)-i-1]
			}
			if curr.children[c-'a'] == nil {
				curr.children[c-'a'] = NewTrieNode()
			}
			curr = curr.children[c-'a']
		}
		curr.leaf = append(curr.leaf, k)
	}
	return root
}

func (trie *Trie) Search(word string, isForward bool) []int {
	curr := trie
	for i := 0; i < len(word); i++ {
		c := byte('a')
		if isForward {
			c = word[i]
		} else {
			c = word[len(word)-i-1]
		}
		if curr.children[c-'a'] == nil {
			return []int{}
		}
		curr = curr.children[c-'a']
	}
	return curr.leaf
}

type WordFilter struct {
	forwardTrie  *Trie
	backwardTrie *Trie
}

func Constructor(words []string) WordFilter {
	ret := WordFilter{}
	ret.forwardTrie = BuildTrie(words, true)
	ret.backwardTrie = BuildTrie(words, false)
	return ret
}

func (this *WordFilter) F(prefix string, suffix string) int {
	forwards := this.forwardTrie.Search(prefix, true)
	backwards := this.backwardTrie.Search(suffix, false)
	// fmt.Println(forwards)
	// fmt.Println(backwards)

	// 1 2 3 6
	// 3 4 5
	a := len(forwards) - 1
	b := len(backwards) - 1
	for a >= 0 && b >= 0 {
		if forwards[a] == backwards[b] {
			return forwards[a]
		}
		if forwards[a] > backwards[b] {
			a--
		} else {
			b--
		}
	}
	return -1
}

func main() {
	filter := Constructor([]string{"pop"})
	fmt.Println(filter.F("", ""))
}
