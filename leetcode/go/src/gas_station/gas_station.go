// 2018-01-15 13:39
package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	tank := -1
	for start < len(gas) && tank < 0 {
		tank = 0
		for i := 0; i < len(gas); i++ {
			index := (start + i) % len(gas)
			tank += gas[index] - cost[index]
			if tank < 0 {
				start += i + 1
				break
			}
		}
	}
	if start > len(gas)-1 {
		return -1
	} else {
		return start
	}
}
func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2}, []int{2, 2}))
}
