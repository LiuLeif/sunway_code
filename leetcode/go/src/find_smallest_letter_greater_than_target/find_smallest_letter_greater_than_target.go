// 2018-01-29 18:44
package main

import (
	"fmt"
	"sort"
)

func binary_search(letters []byte, target byte, f, t int) int {
	for f <= t {
		mid := f + (t-f)/2
		if letters[mid] <= target {
			f = mid + 1
		} else {
			t = mid - 1
		}
	}
	return f
}

func nextGreatestLetter2(letters []byte, target byte) byte {
	index := binary_search(letters, target, 0, len(letters)-1)

	if index == len(letters) {
		return letters[0]
	}
	return letters[index]
}
func nextGreatestLetter(letters []byte, target byte) byte {
	index := sort.Search(len(letters), func(i int) bool {
		return letters[i] > target
	})

	if index == len(letters) {
		return letters[0]
	}
	return letters[index]
}
func main() {
	fmt.Println(string(nextGreatestLetter2([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
}
