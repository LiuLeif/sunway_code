// 2018-02-24 11:44
package main

import (
	"fmt"
	"strings"
)

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func letterCasePermutationRecursive(S string) []string {
	i := 0
	for i < len(S) && !isLetter(S[i]) {
		i++
	}
	if i == len(S) {
		return []string{S}
	}
	prefix := S[:i]
	tmp := letterCasePermutationRecursive(S[i+1:])
	ret := []string{}
	for _, s := range tmp {
		ret = append(ret, prefix+strings.ToLower(string(S[i]))+s)
		ret = append(ret, prefix+strings.ToUpper(string(S[i]))+s)
	}
	return ret
}

func dfs(record, S string, ret *[]string) {
	if len(S) == 0 {
		*ret = append(*ret, record)
		return
	}
	if isLetter(S[0]) {
		dfs(record+strings.ToLower(string(S[0])), S[1:], ret)
		dfs(record+strings.ToUpper(string(S[0])), S[1:], ret)
	} else {
		dfs(record+string(S[0]), S[1:], ret)
	}
}

func letterCasePermutation(S string) []string {
	ret := []string{}
	dfs("", S, &ret)
	return ret
}

func main() {
	fmt.Println(letterCasePermutation("abc"))
	fmt.Println(letterCasePermutationRecursive("abc"))
}
