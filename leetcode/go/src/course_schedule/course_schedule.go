// 2017-11-14 18:54
package main

import "fmt"

func dfs(m map[int]*[]int, visisted map[int]bool, n int) bool {
	if visisted[n] {
		return false
	}
	neighbor := m[n]
	if neighbor == nil {
		return true
	}
	visisted[n] = true
	for _, v := range *neighbor {
		if !dfs(m, visisted, v) {
			return false
		}
	}
	visisted[n] = false
	return true
}

func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	m := make(map[int]*[]int)

	for _, v := range prerequisites {
		a := v[0]
		b := v[1]
		if m[b] == nil {
			m[b] = &[]int{a}
		} else {
			t := m[b]
			*t = append(*t, a)
		}
	}

	visited := make(map[int]bool)
	for i := 0; i < numCourses; i++ {
		if !dfs(m, visited, i) {
			return false
		}
	}

	return true
}

func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	m := make(map[int]*[]int)
	ingress := make(map[int]int)

	for _, v := range prerequisites {
		a := v[0]
		b := v[1]
		ingress[a]++
		if m[b] == nil {
			m[b] = &[]int{a}
		} else {
			t := m[b]
			*t = append(*t, a)
		}
	}

	queue := []int{}
	count := 0
	for i := 0; i < numCourses; i++ {
		if ingress[i] == 0 {
			queue = append(queue, i)
			count++
		}
	}
	fmt.Println(queue)
	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]
		neighbor := m[t]
		if neighbor != nil {
			for _, v := range *neighbor {
				ingress[v]--
				if ingress[v] == 0 {
					queue = append(queue, v)
					count++
				}
			}
		}
	}

	return count == numCourses
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int]*[]int)
	ingress := make(map[int]int)

	for _, v := range prerequisites {
		a := v[0]
		b := v[1]
		ingress[a]++
		if m[b] == nil {
			m[b] = &[]int{a}
		} else {
			t := m[b]
			*t = append(*t, a)
		}
	}

	count := 0
	again := true
	removed := make(map[int]bool)
	for again {
		again = false
		for i := 0; i < numCourses; i++ {
			if removed[i] {
				continue
			}
			if ingress[i] == 0 {
				removed[i] = true
				again = true
				count++
				neighbor := m[i]
				if neighbor != nil {
					for _, v := range *neighbor {
						ingress[v]--
					}
				}
				break
			}
		}
	}

	return count == numCourses
}
func main() {
	prerequisites := make([][]int, 2)
	prerequisites[0] = []int{1, 0}
	prerequisites[1] = []int{0, 1}
	fmt.Println(canFinishDFS(3, prerequisites))
}
