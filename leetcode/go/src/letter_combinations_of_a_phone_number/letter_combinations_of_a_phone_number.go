// 2018-01-02 13:51
package main

import "fmt"

func helper(ret *[]string, record, digits string, m map[byte]string, start int) {
	if start == len(digits) {
		if len(record) != 0 {
			*ret = append(*ret, record)
		}
		return
	}
	digit := byte(digits[start])
	for _, char := range m[digit] {
		helper(ret, record+string(char), digits, m, start+1)
	}
}

func letterCombinations(digits string) []string {
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	ret := []string{}
	helper(&ret, "", digits, m, 0)
	return ret
}
func main() {
	fmt.Println(letterCombinations("23"))
}
