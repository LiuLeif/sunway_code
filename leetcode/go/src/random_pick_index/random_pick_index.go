// 2018-02-12 11:34
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	ret := Solution{}
	ret.nums = nums
	return ret
}

func (this *Solution) Pick(target int) int {
	ret := -1
	n := 0
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] == target {
			n++
			if rand.Intn(n) == 0 {
				ret = i
			}
		}
	}
	return ret
}

func main() {
	rand.Seed(time.Now().Unix())
	obj := Constructor([]int{1, 2, 3, 3, 3})
	fmt.Println(obj.Pick(3))
}
