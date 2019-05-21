// 2018-01-27 18:33
package main

import (
	"fmt"
)

type Node struct {
	children [26]*Node
	val      int
	leaf     bool
}
type MapSum struct {
	root *Node
}

func Constructor() MapSum {
	return MapSum{&Node{}}
}

func (this *MapSum) Insert(key string, val int) {
	root := this.root
	for i := 0; i < len(key); i++ {
		if root.children[key[i]-'a'] == nil {
			root.children[key[i]-'a'] = &Node{}
		}
		root = root.children[key[i]-'a']
	}
	root.leaf = true
	root.val = val
}

func sum(root *Node) int {
	ret := 0
	if root.leaf {
		ret = root.val
	}
	for i := 0; i < 26; i++ {
		if root.children[i] != nil {
			ret += sum(root.children[i])
		}
	}
	return ret
}

func (this *MapSum) Sum(prefix string) int {
	root := this.root

	for i := 0; i < len(prefix); i++ {
		c := prefix[i] - 'a'
		if root.children[c] == nil {
			return 0
		}
		root = root.children[c]
	}

	return sum(root)
}

func main() {
	obj := Constructor()
	obj.Insert("hello", 1)
	obj.Insert("hell", 2)
	obj.Insert("hell", 1)
	fmt.Println(obj.Sum("hell"))
}
