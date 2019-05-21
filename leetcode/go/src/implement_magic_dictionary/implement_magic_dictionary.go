// 2018-03-30 12:38
package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isLeaf   bool
}

func NewTrie() *Trie {
	return &Trie{}
}

func (trie *Trie) Add(word string) {
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if trie.children[c] == nil {
			trie.children[c] = NewTrie()
		}
		trie = trie.children[c]
	}
	trie.isLeaf = true
}

func (trie *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if trie.children[c] == nil {
			return false
		}
		trie = trie.children[c]
	}
	return trie.isLeaf
}

type MagicDictionary struct {
	trie *Trie
}

func Constructor() MagicDictionary {
	return MagicDictionary{NewTrie()}
}

func (this *MagicDictionary) BuildDict(dict []string) {
	for _, v := range dict {
		this.trie.Add(v)
	}
}

func (this *MagicDictionary) search(trie *Trie, word string) bool {
	if len(word) == 0 {
		return false
	}
	for i := 0; i < 26; i++ {
		child := trie.children[i]
		if child == nil {
			continue
		}
		ret := false
		if i == int(word[0]-'a') {
			ret = this.search(child, word[1:])
		} else {
			ret = child.Search(word[1:])
		}
		if ret {
			return true
		}
	}
	return false
}

func (this *MagicDictionary) Search(word string) bool {
	return this.search(this.trie, word)
}

func main() {
	dict := Constructor()
	dict.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(dict.Search("hello"))
	fmt.Println(dict.Search("hhllo"))
}
