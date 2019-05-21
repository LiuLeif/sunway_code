// 2017-12-27 17:05
package main

import (
	"fmt"
	"math"
)

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

type TrieNode struct {
	children [2]*TrieNode
	value    int
}

func buildTrie(nums []int) *TrieNode {
	root := &TrieNode{}
	for i := 0; i < len(nums); i++ {
		currNode := root
		for j := 31; j >= 0; j-- {
			bit := (nums[i] >> uint(j)) & 1
			if currNode.children[bit] == nil {
				currNode.children[bit] = &TrieNode{}
			}
			currNode = currNode.children[bit]
		}
		currNode.value = nums[i]
	}
	return root
}

func findMaximumXOR(nums []int) int {
	root := buildTrie(nums)
	ret := 0
	for i := 0; i < len(nums); i++ {
		currNode := root
		for j := 31; j >= 0; j-- {
			bit := (nums[i] >> uint(j)) & 1
			x := bit ^ 1
			if currNode.children[x] != nil {
				currNode = currNode.children[x]
			} else {
				currNode = currNode.children[^x&1]
			}
		}
		ret = max(ret, nums[i]^currNode.value)
	}
	return ret
}
func main() {
	// 8   1000
	// 10  1010
	// 2   0010
	// xor 8^2=1010=10
	fmt.Println(findMaximumXOR([]int{8, 10, 2}))
}
