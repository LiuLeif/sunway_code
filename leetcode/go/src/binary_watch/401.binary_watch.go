// 2018-11-04 12:18
package main

import "fmt"

func readBinaryWatch(num int) []string {
	ret := []string{}
	// 0~11 1011
	// 0~59
	to_string := func(n int) string {
		hour := (n >> 6) & 0xf
		min := n & 0x3f
		if hour < 12 && min < 60 {
			return fmt.Sprintf("%d:%02d", hour, min)
		}
		return ""
	}
	var dfs func(curr, n, index int)
	dfs = func(curr, n, index int) {
		if n == 0 {
			t := to_string(curr)
			if len(t) != 0 {
				ret = append(ret, t)
			}
			return
		}
		for i := index; i < 10; i++ {
			dfs(curr|(1<<uint(i)), n-1, i+1)
		}
	}
	dfs(0, num, 0)
	return ret
}

func main() {
	fmt.Println(readBinaryWatch(2))
}
