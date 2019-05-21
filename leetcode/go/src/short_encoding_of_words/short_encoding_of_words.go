// 2018-10-18 14:25
package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	len      int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func (this *TrieNode) Insert(S string) {
	root := this
	for i := len(S) - 1; i >= 0; i-- {
		c := S[i]
		if root.children[c-97] == nil {
			root.children[c-97] = &TrieNode{}
		}
		root = root.children[c-97]
	}
	root.len = len(S)
}

func (this *TrieNode) FindLeaf() int {
	if this == nil {
		return 0
	}

	ret := 0
	for i := 0; i < 26; i++ {
		if this.children[i] != nil {
			ret += this.children[i].FindLeaf()
		}
	}
	if ret == 0 {
		ret = this.len + 1
	}
	return ret
}

func minimumLengthEncoding(words []string) int {
	trieNode := NewTrieNode()
	for _, w := range words {
		trieNode.Insert(w)
	}
	return trieNode.FindLeaf()
}
func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell", "imea"}))
}
