// 2017-11-29 15:31
package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}

	retSet := make(map[string]bool)
	hashSeen := make(map[int]bool)
	index := map[byte]int{'A': 0, 'T': 1, 'C': 2, 'G': 3}

	hash := 0
	for i := 0; i < 10; i++ {
		x := index[s[i]] << uint(18-i*2)
		hash |= x
	}

	hashSeen[hash] = true

	for i := 1; i < len(s)-9; i++ {
		hash = hash << 2
		hash |= index[s[i+9]]
		hash &= (1 << 20) - 1
		if hashSeen[hash] {
			retSet[s[i:i+10]] = true
		}
		hashSeen[hash] = true
	}

	ret := []string{}
	for k, _ := range retSet {
		ret = append(ret, k)
	}
	return ret
}
func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAA"))
}
