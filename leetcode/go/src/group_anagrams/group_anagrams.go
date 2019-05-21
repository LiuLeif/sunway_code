// 2017-11-28 18:55
package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 107}

	m := make(map[int][]string)
	for _, s := range strs {
		hash := 1
		for i := 0; i < len(s); i++ {
			hash *= primes[int(s[i]-'a')]
		}
		if m[hash] == nil {
			m[hash] = []string{}
		}
		m[hash] = append(m[hash], s)
	}

	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}
func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
