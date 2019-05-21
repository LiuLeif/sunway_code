// 2017-11-23 10:39
package main

import "fmt"

type Trie struct {
	Next map[byte]*Trie
	Word string
}

/** Initialize your data structure here. */
func Constructor() Trie {
	ret := Trie{}
	ret.Next = make(map[byte]*Trie)
	return ret
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		if root.Next[word[i]] == nil {
			trie := Constructor()
			root.Next[word[i]] = &trie
		}
		root = root.Next[word[i]]
	}
	root.Word = word
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for i := 0; i < len(word); i++ {
		root = root.Next[word[i]]
		if root == nil {
			return false
		}
	}
	return root.Word != ""
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for i := 0; i < len(prefix); i++ {
		root = root.Next[prefix[i]]
		if root == nil {
			return false
		}
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("hello world")
	fmt.Println(trie.Search("hello"))
	fmt.Println(trie.StartsWith("hello"))
}
