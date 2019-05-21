// 2018-01-24 16:59
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	loc  map[int]int
	nums []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.loc[val]; ok {
		return false
	}
	this.loc[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.loc[val]; !ok {
		return false
	}

	thisIndex := this.loc[val]
	lastVal := this.nums[len(this.nums)-1]
	this.nums[thisIndex] = lastVal
	this.loc[lastVal] = thisIndex

	delete(this.loc, val)
	this.nums = this.nums[:len(this.nums)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

func main() {
	rand.Seed(time.Now().Unix())
	obj := Constructor()
	obj.Insert(1)
	obj.Insert(2)
	obj.Insert(3)
	obj.Insert(4)
	obj.Insert(5)
	fmt.Println(obj.GetRandom())
	obj.Remove(1)
	fmt.Println(obj.GetRandom())
}
