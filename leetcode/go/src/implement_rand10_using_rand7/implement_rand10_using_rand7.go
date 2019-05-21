// 2018-10-18 16:57
package main

import "fmt"
import "math/rand"
import "time"

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	result := 40
	for result >= 40 {
		result = 7*(rand7()-1) + rand7() - 1
	}
	return result%10 + 1
}
func main() {
	rand.Seed(int64(time.Now().Second()))
	for i := 0; i < 10; i++ {
		fmt.Println(rand10())
	}
}
