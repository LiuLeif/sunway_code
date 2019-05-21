// 2017-12-12 16:01
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func nextGreaterElement(n int) int {
	nums := strings.Split(strconv.Itoa(n), "")
	start := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			start = i - 1
			break
		}
	}
	if start == -1 {
		return -1
	}
	end := -1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[start] {
			end = i
			break
		}
	}
	nums[start], nums[end] = nums[end], nums[start]
	// reverse nums[start+1:]
	tmp := nums[start+1:]
	for i := 0; i < len(tmp)/2; i++ {
		tmp[i], tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1], tmp[i]
	}
	ret, _ := strconv.Atoi(strings.Join(nums, ""))
	if ret > math.MaxInt32 {
		return -1
	}
	return ret
}

func main() {
	fmt.Println(nextGreaterElement(1999999999))
}
