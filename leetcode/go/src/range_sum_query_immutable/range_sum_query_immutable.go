// 2018-01-04 17:04
package main

import "fmt"

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	ret := NumArray{}
	ret.sum = make([]int, len(nums)+1)
	if len(nums) == 0 {
		return ret
	}
	for i := 0; i < len(nums); i++ {
		ret.sum[i+1] = ret.sum[i] + nums[i]
	}
	return ret
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum[j+1] - this.sum[i]
}

func main() {
	array := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(array.SumRange(0, 2))
	fmt.Println(array.SumRange(2, 5))
	fmt.Println(array.SumRange(0, 5))
}
