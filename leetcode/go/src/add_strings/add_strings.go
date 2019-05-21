// 2018-01-31 13:31
package main

import (
	"fmt"
	"math"
	"strconv"
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

func swap(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

func addStrings(num1 string, num2 string) string {
	x1, x2 := make([]int, len(num1)), make([]int, len(num2))
	for i := 0; i < len(num1); i++ {
		x1[i] = int(num1[i] - '0')
	}
	for i := 0; i < len(num2); i++ {
		x2[i] = int(num2[i] - '0')
	}

	swap(x1)
	swap(x2)

	result := []int{}
	carry := 0
	for i := 0; i < max(len(x1), len(x2)); i++ {
		a, b := 0, 0
		if i < len(x1) {
			a = x1[i]
		}
		if i < len(x2) {
			b = x2[i]
		}
		sum := a + b + carry
		result = append(result, sum%10)
		carry = sum / 10
	}
	if carry != 0 {
		result = append(result, carry)
	}
	swap(result)
	ret := ""
	for i := 0; i < len(result); i++ {
		ret += strconv.Itoa(result[i])
	}
	return ret
}
func main() {
	fmt.Println(addStrings("123", "7"))
}
