// 2017-12-25 17:35
package main

import "fmt"

func isCandidate(x, y string) bool {
	count := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			count++
		}
	}
	return count == 1
}

func minMutation(start string, end string, bank []string) int {
	queue := []string{start, ""}
	count := 0
	visited := make(map[string]bool)
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == "" {
			count++
			queue = append(queue, "")
			continue
		}
		if top == end {
			return count
		}
		for i := 0; i < len(bank); i++ {
			if !visited[bank[i]] {
				if isCandidate(top, bank[i]) {
					queue = append(queue, bank[i])
					visited[bank[i]] = true
				}
			}
		}
	}

	return -1
}
func main() {
	start := "AACCGGTT"
	end := "AAACGGTA"
	bank := []string{
		"AACCGGTA", "AACCGCTA", "AAACGGTA",
	}
	fmt.Println(minMutation(start, end, bank))
}
