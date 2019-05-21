// 2018-01-12 13:00
package main

import (
	"fmt"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	// binary search 的原理:
	//
	// 1 2 3 4 5 6 7 8, k=3, x=4 时, 计算各个点 与 x 的差的 abs:
	// 3 2 1 0 1 2 3 4, 再计算 [i,i+k) 的和:
	// 6 3 2 3 6 9
	//
	// 则原问题等价为: 对于 6 3 2 3 6 9, 查找其最小值 (即 2).
	// 这个序列不是单调的, 但可以用 binary search 找到最小值:
	//
	// 例如:
	// 6 3 2 3 6 9, mid 为 3 时, arr[mid] <arr[mid+1], 所以最小值必然在 [f,mid]
	//       ^
	// 若 mid 为 1, arr[mid] > arr[mid+1], 最小值在 (mid, t]
	// 而比较 x[mid],x[mid+1] 即是比较 x-arr[mid] 与 arr[mid+k]-x

	lo, hi := 0, len(arr)-k
	for lo < hi {
		mid := (lo + hi) / 2
		if x-arr[mid] > arr[mid+k]-x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return arr[lo : lo+k]
}

func firstLargerOrEqual(arr []int, x int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func findClosestElements2(arr []int, k int, x int) []int {
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= x
	})
	i, j := index-1, index

	for k > 0 {
		if i < 0 || (j < len(arr) && x-arr[i] > arr[j]-x) {
			j++
		} else {
			i--
		}
		k--
	}

	return arr[i+1 : j]
}
func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 2, 6))
	fmt.Println(findClosestElements2([]int{1, 2, 3, 4, 5}, 2, 6))
}
