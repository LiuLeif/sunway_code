// 2018-10-31 21:41
package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		x := 0
		i++
		for i != len(name) && name[i] == name[i-1] {
			i++
			x++
		}
		y := 0
		j++
		for j != len(typed) && typed[j] == typed[j-1] {
			j++
			y++
		}
		if y < x {
			return false
		}
	}
	return i == len(name) && j == len(typed)
}

func main() {
	fmt.Println(isLongPressedName("pyplrz", "ppyypllr"))
}
