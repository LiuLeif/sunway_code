// 2017-11-27 14:25
package main

import "fmt"

func totalHammingDistance(nums []int) int {
	ret := 0
	for j := 0; j < 32; j++ {
		bitSet := 0
		for i := 0; i < len(nums); i++ {
			if nums[i]&(1<<uint(j)) != 0 {
				bitSet++
			}
		}
		ret += (len(nums) - bitSet) * bitSet
	}

	return ret
}

func main() {
	//   100
	//  1110
	//    10

	fmt.Println(totalHammingDistance([]int{4, 14, 2}))
}
