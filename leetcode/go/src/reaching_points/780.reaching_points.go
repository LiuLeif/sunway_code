// 2018-11-06 17:14
package main

import "fmt"

type Pair struct {
	a, b int
}

func reachingPoints2(sx int, sy int, tx int, ty int) bool {
	queue := []Pair{{sx, sy}}
	visited := map[Pair]bool{Pair{sx, sy}: true}
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if top.a == tx && top.b == ty {
			return true
		}

		t1 := Pair{top.a + top.b, top.b}
		if !visited[t1] && (t1.a <= tx && t1.b <= ty) {
			visited[t1] = true
			queue = append(queue, t1)
		}

		t2 := Pair{top.a, top.a + top.b}
		if !visited[t2] && (t2.a <= tx && t2.b <= ty) {
			visited[t2] = true
			queue = append(queue, t2)
		}
	}
	return false
}

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx > sx && ty > sy {
		tx, ty = tx%ty, ty%tx
		// if tx > ty {
		// 	tx, ty = tx%ty, ty
		// } else {
		// 	tx, ty = tx, ty%tx
		// }
	}
	return (tx == sx && (ty-sy)%sx == 0) || (ty == sy && (tx-sx)%sy == 0)
}

func main() {
	fmt.Println(reachingPoints(1, 1, 3, 5))
	fmt.Println(reachingPoints(35, 13, 455955547, 455955547))
}
