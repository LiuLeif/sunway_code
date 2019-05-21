// 2018-04-08 14:16
package main

import "fmt"

func max(nums ...float64) float64 {
	ret := 0.0
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func largestSumOfAverages(A []int, K int) float64 {
	sum := make([]int, len(A)+1)
	// sum[a,b)=total(b)-total(a)
	for i := 1; i <= len(A); i++ {
		sum[i] = sum[i-1] + A[i-1]
	}
	mem := map[[2]int]float64{}
	var check func(int, int) float64
	check = func(f int, k int) float64 {
		if cache, ok := mem[[2]int{f, k}]; ok {
			return cache
		}
		if k == 1 {
			ret := float64(sum[len(A)]-sum[f]) / float64(len(A)-f)
			mem[[2]int{f, k}] = ret
			return ret
		}
		ret := 0.0
		for i := f; i <= len(A)-k; i++ {
			avg := float64(sum[i+1]-sum[f]) / float64(i-f+1)
			ret = max(ret, avg+check(i+1, k-1))
		}
		mem[[2]int{f, k}] = ret
		return ret
	}
	return check(0, K)
}
func main() {
	// 9   1 2 3   9
	fmt.Println(largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
}
