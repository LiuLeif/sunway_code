// 2018-10-19 11:16
package main

import "fmt"

func pushDominoes(dominoes string) string {
	prevLeft, prevRight := -1, -1
	v := []rune(dominoes)
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i] == '.' {
			continue
		}
		if dominoes[i] == 'L' {
			if prevRight != -1 {
				for a, b := prevRight+1, i-1; a < b; a, b = a+1, b-1 {
					v[a] = 'R'
					v[b] = 'L'
				}
			} else {
				for j := prevLeft + 1; j < i; j++ {
					// 0 1 2 3 4
					v[j] = 'L'
				}
			}
			prevLeft = i
			prevRight = -1
		} else {
			if prevRight != -1 {
				for j := prevRight + 1; j < i; j++ {
					v[j] = 'R'
				}
			}
			prevRight = i
			prevLeft = -1
		}
	}
	if prevRight != -1 {
		for i := prevRight + 1; i < len(v); i++ {
			v[i] = 'R'
		}
	}
	return string(v)
}
func main() {
	fmt.Println(pushDominoes(".RL"))
}
