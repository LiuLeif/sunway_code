// 2017-11-23 10:39
package main

import "fmt"

type WordDictionary struct {
	Next map[byte]*WordDictionary
	Word string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	ret := WordDictionary{}
	ret.Next = make(map[byte]*WordDictionary)
	return ret
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		if root.Next[word[i]] == nil {
			t := Constructor()
			root.Next[word[i]] = &t
		}
		root = root.Next[word[i]]
	}
	root.Word = word
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	root := this
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			for _, v := range root.Next {
				if v.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
		if root.Next[word[i]] == nil {
			return false
		}
		root = root.Next[word[i]]
	}
	return root.Word != ""
}

func main() {
	dict := Constructor()
	dict.AddWord("hello")
	dict.AddWord("ab")
	fmt.Println(dict.Search("hello"))
	fmt.Println(dict.Search("hell."))
	fmt.Println(dict.Search(".."))
}
