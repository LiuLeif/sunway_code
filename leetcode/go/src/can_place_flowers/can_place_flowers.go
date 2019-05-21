// 2017-11-14 18:54
package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}
		left := i == 0 || (flowerbed[i-1] == 0)
		right := i == len(flowerbed)-1 || (flowerbed[i+1] == 0)
		if left && right {
			n -= 1
			flowerbed[i] = 1
		}
	}

	return n <= 0
}
func main() {
	fmt.Println(canPlaceFlowers([]int{0}, 1))
}
