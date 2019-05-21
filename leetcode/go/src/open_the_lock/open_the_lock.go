// 2018-01-04 17:04
package main

import "fmt"

type Quad [4]byte

func NewQuad(s string) Quad {
	ret := Quad{}
	copy(ret[:], []byte(s))
	return ret
}

func DupQuad(orig Quad) Quad {
	ret := Quad{}
	copy(ret[:], (orig)[:])
	return ret
}

func openLock(deadends []string, target string) int {
	m := map[Quad]bool{}
	for i := 0; i < len(deadends); i++ {
		m[NewQuad(deadends[i])] = true
	}
	nilQuad := NewQuad("aaaa")
	queue := []Quad{NewQuad("0000"), nilQuad}
	ret := 0
	targetQuad := NewQuad(target)
	if m[NewQuad("0000")] {
		return -1
	}
	for len(queue) != 1 {
		top := queue[0]
		queue = queue[1:]
		if top == nilQuad {
			ret++
			queue = append(queue, nilQuad)
			continue
		}
		if top == targetQuad {
			return ret
		}
		for i := 0; i < 4; i++ {
			q := DupQuad(top)
			q[i] = (q[i]+1-48)%10 + 48
			if !m[q] {
				queue = append(queue, q)
				m[q] = true
			}
			q = DupQuad(top)
			q[i] = (q[i]+9-48)%10 + 48
			if !m[q] {
				queue = append(queue, q)
				m[q] = true
			}
		}
	}
	return -1
}
func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"8888"}, "0009"))
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
	fmt.Println(openLock([]string{"0000"}, "8888"))
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
