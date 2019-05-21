// 2018-01-24 15:29
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	origNums []int
	currNums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums, append([]int{}, nums...)}
}

func (this *Solution) Reset() []int {
	return this.origNums
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.currNums); i++ {
		t := rand.Intn(len(this.currNums)-i) + i
		this.currNums[i], this.currNums[t] = this.currNums[t], this.currNums[i]
	}
	return this.currNums
}

func main() {
	rand.Seed(time.Now().Unix())
	obj := Constructor([]int{1, 2, 3, 4, 5})
	obj.Reset()
	fmt.Println(obj.Shuffle())
	fmt.Println(obj.Shuffle())
}
