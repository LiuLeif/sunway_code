// 2018-01-12 15:24
package main

import "fmt"

func getHint(secret string, guess string) string {
	bull, cow := 0, 0
	ma := map[byte]int{}
	mb := map[byte]int{}
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bull++
		} else {
			if ma[guess[i]] != 0 {
				cow++
				ma[guess[i]]--
			} else {
				mb[guess[i]]++
			}
			if mb[secret[i]] != 0 {
				cow++
				mb[secret[i]]--
			} else {
				ma[secret[i]]++
			}
		}
	}
	return fmt.Sprintf("%dA%dB", bull, cow)
}
func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1234", "1124"))
}
