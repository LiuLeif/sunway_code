// 2017-11-24 14:50
package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	m := make(map[int][]int)
	ingress := make(map[int]int)

	for _, v := range prerequisites {
		a := v[0]
		b := v[1]
		ingress[a]++
		if m[b] == nil {
			m[b] = []int{a}
		} else {
			m[b] = append(m[b], a)
		}
	}

	ret := []int{}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if ingress[i] == 0 {
			ret = append(ret, i)
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]
		neighbor := m[t]
		if neighbor != nil {
			for _, v := range neighbor {
				ingress[v]--
				if ingress[v] == 0 {
					ret = append(ret, v)
					queue = append(queue, v)
				}
			}
		}
	}

	if len(ret) != numCourses {
		return []int{}
	}

	return ret
}
func main() {
	fmt.Println(findOrder(
		4, [][]int{
			{0, 1},
			{2, 0},
			{3, 1},
			{3, 2},
		},
	))
}
