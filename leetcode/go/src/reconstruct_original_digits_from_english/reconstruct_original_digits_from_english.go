// 2018-01-05 11:30
package main

import "fmt"

func remove(s string, n int, count []int) {
	for i := 0; i < len(s); i++ {
		count[s[i]-'a'] -= n
	}
}

func originalDigits(s string) string {
	count := [26]int{}
	mark := [10]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	mark[0] = count['z'-'a']
	remove("zero", mark[0], count[:])

	mark[6] = count['x'-'a']
	remove("six", mark[6], count[:])

	mark[2] = count['w'-'a']
	remove("two", mark[2], count[:])

	mark[8] = count['g'-'a']
	remove("eight", mark[8], count[:])

	mark[4] = count['u'-'a']
	remove("four", mark[4], count[:])

	mark[3] = count['r'-'a']
	remove("three", mark[3], count[:])

	mark[5] = count['f'-'a']
	remove("five", mark[5], count[:])

	mark[7] = count['s'-'a']
	remove("seven", mark[7], count[:])

	mark[9] = count['i'-'a']
	remove("nine", mark[9], count[:])

	mark[1] = count['o'-'a']
	remove("one", mark[1], count[:])

	ret := []byte{}
	for i := 0; i < 10; i++ {
		for j := 0; j < mark[i]; j++ {
			ret = append(ret, byte(i+48))
		}
	}

	return string(ret)
}
func main() {
	fmt.Println(originalDigits("twtohree"))
}
