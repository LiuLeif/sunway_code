// 2018-10-24 09:21
package main

import "fmt"
import "strconv"
import "strings"

func optimalDivision(nums []int) string {
	ret := strconv.Itoa(nums[0])
	if len(nums) == 1 {
		return ret
	}
	ret += "/"

	numStr := make([]string, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		numStr[i-1] = strconv.Itoa(nums[i])
	}
	if len(nums) == 2 {
		ret += strings.Join(numStr, "/")
	} else {
		ret += "(" + strings.Join(numStr, "/") + ")"
	}

	return ret
}

func main() {
	fmt.Println(optimalDivision([]int{1000, 1, 2}))
}
