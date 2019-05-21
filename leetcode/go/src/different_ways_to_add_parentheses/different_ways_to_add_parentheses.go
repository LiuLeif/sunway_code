// 2018-01-19 10:52
package main

import "fmt"

func tokenize(s string) ([]int, []byte) {
	nums := []int{}
	ops := []byte{}
	digit := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+', '-', '*':
			nums = append(nums, digit)
			digit = 0
			ops = append(ops, s[i])
		default:
			digit = digit*10 + int(s[i]-'0')
		}
	}
	nums = append(nums, digit)
	return nums, ops
}

func helper(nums []int, ops []byte) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	ret := []int{}

	for i := 0; i < len(nums)-1; i++ {
		op := ops[i]
		left := helper(nums[:i+1], ops[:i])
		right := helper(nums[i+1:], ops[i+1:])
		for _, l := range left {
			for _, r := range right {
				result := 0
				switch op {
				case '+':
					result = l + r
				case '-':
					result = l - r
				case '*':
					result = l * r
				}
				ret = append(ret, result)
			}
		}
	}
	return ret
}

func diffWaysToCompute(input string) []int {
	return helper(tokenize(input))
}
func main() {
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
