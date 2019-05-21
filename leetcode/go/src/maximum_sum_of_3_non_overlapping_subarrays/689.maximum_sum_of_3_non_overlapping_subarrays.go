// 2018-12-05 10:11
package main

import "fmt"

func maxSumOfThreeSubarrays(nums []int, K int) []int {
	// 1 2 1 2 6 7 5 1
	// 0 1 2 3 4 5 6 7
	// i     j     k
	accu := make([]int, len(nums)+1)
	sum := 0
	for i := 1; i < len(nums)+1; i++ {
		sum += nums[i-1]
		accu[i] = sum
	}

	leftV := make([]int, len(nums))
	leftI := make([]int, len(nums))
	bestV, bestI := 0, 0
	for i := K - 1; i < len(nums); i++ {
		tmp := accu[i+1] - accu[i-K+1]
		if tmp > bestV {
			bestV = tmp
			bestI = i - K + 1
		}
		leftV[i] = bestV
		leftI[i] = bestI
	}

	rightV := make([]int, len(nums))
	rightI := make([]int, len(nums))
	bestV, bestI = 0, 0
	for i := len(nums) - K; i >= 0; i-- {
		tmp := accu[i+K] - accu[i]
		if tmp >= bestV {
			bestV = tmp
			bestI = i
		}
		rightV[i] = bestV
		rightI[i] = bestI
	}

	ret := []int{}
	bestV = 0
	for i := K; i < len(nums)-2*K+1; i++ {
		tmp := leftV[i-1] + rightV[i+K] + accu[i+K] - accu[i]
		if tmp > bestV {
			bestV = tmp
			ret = []int{leftI[i-1], i, rightI[i+K]}
		}
	}
	return ret
}

func main() {
	fmt.Println(maxSumOfThreeSubarrays([]int{7, 13, 20, 19, 19, 2, 10, 1, 1, 19}, 3))
	fmt.Println(maxSumOfThreeSubarrays([]int{1, 2, 1, 2, 6, 7, 5, 18}, 2))
}
