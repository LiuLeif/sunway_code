// 2017-11-14 18:54
package main

import (
	"fmt"
	"strconv"
)

func max(nums ...byte) byte {
	ret := byte(0)
	for i := 0; i < len(nums); i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
	}
	return ret
}

func maximumSwap(num int) int {
	s := []rune(strconv.Itoa(num))
	// 2736
	for i := 0; i < len(s)-1; i++ {
		maxValue := s[i]
		maxIndex := i
		for j := len(s) - 1; j >= i+1; j-- {
			if s[j] > maxValue {
				maxValue = s[j]
				maxIndex = j
			}
		}

		if maxIndex != i {
			s[maxIndex], s[i] = s[i], s[maxIndex]
			break
		}
	}

	ret, _ := strconv.Atoi(string(s))
	return ret
}
func maximumSwap2(num int) int {
	x := []byte(strconv.Itoa(num))
	y := make([]byte, len(x))
	index := make([]int, len(x))
	index[len(index)-1] = len(x) - 1
	for i := len(x) - 2; i >= 0; i-- {
		y[i] = max(y[i+1], x[i+1])
		if y[i] == y[i+1] {
			index[i] = index[i+1]
		} else {
			index[i] = i + 1
		}
	}
	for i := 0; i < len(x)-1; i++ {
		if y[i] > x[i] {
			x[i], x[index[i]] = x[index[i]], x[i]
			break
		}
	}

	ret, _ := strconv.Atoi(string(x))
	return ret
}
func main() {
	fmt.Println(maximumSwap2(13))
	fmt.Println(maximumSwap2(98368))
}
