// 2018-01-17 15:05
package main

import "fmt"

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func plusOne2(digits []int) []int {
	ret := make([]int, len(digits)+1)
	for i := len(digits) - 1; i >= 0; i-- {
		sum := 0
		if i == len(digits)-1 {
			sum = digits[i] + 1 + ret[i+1]
		} else {
			sum = digits[i] + ret[i+1]
		}
		ret[i+1] = sum % 10
		ret[i] += sum / 10
	}
	if ret[0] == 0 {
		return ret[1:]
	}
	return ret
}

func plusOne(digits []int) []int {
	ret := []int{}
	reverse(digits)
	carry := 0
	plus := []int{1}

	i, j := 0, 0
	for i < len(digits) || j < len(plus) {
		x, y := 0, 0
		if i < len(digits) {
			x = digits[i]
		}
		if j < len(plus) {
			y = plus[j]
		}
		result := x + y + carry
		ret = append(ret, result%10)
		carry = result / 10
		i++
		j++
	}
	if carry != 0 {
		ret = append(ret, carry)
	}
	reverse(ret)
	return ret
}
func main() {
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
	fmt.Println(plusOne2([]int{9, 9, 9, 9}))
}
