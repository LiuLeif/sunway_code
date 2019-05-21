// 2018-11-01 09:42
package main

import (
	"fmt"
)

const (
	HASH_BITS = 5
)

type MyHashSet struct {
	bucket [1 << HASH_BITS][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	ret := MyHashSet{}
	return ret
}

func hash(x int) int {
	return int(uint32(x*2654435761) >> uint32(32-HASH_BITS))

}

func (this *MyHashSet) Add(key int) {
	hashedKey := hash(key)
	if this.bucket[hashedKey] == nil {
		this.bucket[hashedKey] = []int{}
	}
	for i := 0; i < len(this.bucket[hashedKey]); i++ {
		if this.bucket[hashedKey][i] == key {
			return
		}
	}
	this.bucket[hashedKey] = append(this.bucket[hashedKey], key)
}

func (this *MyHashSet) Remove(key int) {
	hashedKey := hash(key)
	if this.bucket[hashedKey] == nil {
		return
	}
	v := this.bucket[hashedKey]
	for i := 0; i < len(v); i++ {
		if v[i] == key {
			this.bucket[hashedKey] = append(this.bucket[hashedKey][:i], this.bucket[hashedKey][i+1:]...)
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	hashedKey := hash(key)
	if this.bucket[hashedKey] == nil {
		return false
	}
	v := this.bucket[hashedKey]
	for i := 0; i < len(v); i++ {
		if v[i] == key {
			return true
		}
	}
	return false
}

func main() {
	hashSet := Constructor()
	hashSet.Add(1)
	hashSet.Add(2)
	fmt.Println(hashSet.Contains(1))
	fmt.Println(hashSet.Contains(3))
	hashSet.Add(2)
	fmt.Println(hashSet.Contains(2))
	hashSet.Remove(2)
	fmt.Println(hashSet.Contains(2))
}
