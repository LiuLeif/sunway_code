// 2018-02-12 15:23
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	ret := make([]int, n)

	currTask := -1
	currStart := -1
	stack := []int{}

	for _, s := range logs {
		m := strings.Split(s, ":")
		task, _ := strconv.Atoi(m[0])
		time, _ := strconv.Atoi(m[2])
		if m[1] == "start" {
			if currTask != -1 {
				ret[currTask] += time - currStart
			}
			currStart = time
			stack = append(stack, currTask)
			currTask = task
		} else {
			ret[currTask] += time - currStart + 1
			currTask = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			currStart = time + 1
		}
	}
	return ret
}
func main() {
	fmt.Println(exclusiveTime(2, []string{
		"0:start:0",
		"1:start:2",
		"1:end:5",
		"0:end:6",
	}))
}
