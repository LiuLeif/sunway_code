// 2018-02-26 11:27
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

func lengthLongestPath(input string) int {
	input = input + string('\n')
	ret := 1
	stack := []int{}
	start := 0
	level := 0
	totalLen := 0
	isFile := false
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '\n':
			wordLen := i - start
			for len(stack) > level {
				totalLen -= stack[len(stack)-1] + 1
				stack = stack[:len(stack)-1]
			}
			totalLen += wordLen + 1
			stack = append(stack, wordLen)

			level = 0
			start = -1
			if isFile {
				ret = max(ret, totalLen)
			}
			isFile = false
		case '\t':
			level++
		case '.':
			isFile = true
		default:
			if start == -1 {
				start = i
			}
		}
	}
	return ret - 1
}
func main() {
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2ext"))
}
