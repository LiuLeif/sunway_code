// 2018-10-30 09:07
package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	first, second int
}

func carFleet(target int, position []int, speed []int) int {
	// 6 8
	// 3 2
	// 1.1  1

	//12 3 7 1  1
	data := make([]Pair, len(position))
	for i := 0; i < len(position); i++ {
		data[i] = Pair{position[i], speed[i]}
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].first < data[j].first
	})

	fleet := 0
	remainingTime := 0.0
	for i := len(data) - 1; i >= 0; i-- {
		time := float64(target-data[i].first) / float64(data[i].second)
		if i == len(data)-1 {
			remainingTime = time
			continue
		}
		if time > remainingTime {
			remainingTime = time
			fleet++
		}
	}

	if remainingTime != 0.0 {
		fleet++
	}
	return fleet
}

func main() {
	fmt.Println(carFleet(10, []int{6, 8}, []int{3, 2}))
}
