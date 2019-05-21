// 2018-02-24 14:35
package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return i == len(bits)-1
}
func isOneBitCharacterRecursive(bits []int) bool {
	if len(bits) == 0 {
		return false
	}
	if len(bits) == 1 {
		return bits[0] == 0
	}
	if bits[0] == 1 {
		return isOneBitCharacterRecursive(bits[2:])
	}
	return isOneBitCharacterRecursive(bits[1:])
}
func main() {
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))
	fmt.Println(isOneBitCharacterRecursive([]int{1, 1, 1, 0}))
}
