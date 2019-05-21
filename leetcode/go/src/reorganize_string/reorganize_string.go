// 2018-01-30 10:13
package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	count int
	char  byte
}

type Heap []Pair

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (ph *Heap) Pop() interface{} {
	t := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]
	return t
}

func (ph *Heap) Push(x interface{}) {
	*ph = append(*ph, x.(Pair))
}

func reorganizeString(S string) string {
	m := map[byte]int{}
	for i := 0; i < len(S); i++ {
		m[S[i]]++
	}

	pq := Heap{}
	for k, v := range m {
		pq = append(pq, Pair{v, k})
	}

	ret := ""
	heap.Init(&pq)
	for len(pq) != 0 {
		a := heap.Pop(&pq).(Pair)
		a.count--
		ret += string(a.char)
		if len(pq) == 0 {
			if a.count > 0 {
				return ""
			} else {
				return ret
			}
		}
		b := heap.Pop(&pq).(Pair)
		b.count--
		ret += string(b.char)
		if a.count != 0 {
			heap.Push(&pq, a)
		}
		if b.count != 0 {
			heap.Push(&pq, b)
		}
	}
	return ret
}
func main() {
	fmt.Println(reorganizeString("abaac"))
}
