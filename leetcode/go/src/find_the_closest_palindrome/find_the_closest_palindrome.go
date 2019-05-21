// 2018-02-05 10:33
package main

import (
	"fmt"
	"strconv"
)

func largerPalindromic(s string) string {
	if isParlindromic(s) {
		return s
	}
	tmp := []byte(s)
	for i := 0; i < len(tmp)/2; i++ {
		tmp[len(tmp)-1-i] = tmp[i]
	}
	t := string(tmp)
	sValue, _ := strconv.Atoi(s)
	tValue, _ := strconv.Atoi(t)
	if sValue < tValue {
		return t
	}
	for i := (len(tmp) - 1) / 2; i >= 0; i-- {
		if tmp[i] != '9' {
			tmp[i] += 1
			break
		} else {
			tmp[i] = '0'
		}
	}
	for i := 0; i < len(tmp)/2; i++ {
		tmp[len(tmp)-1-i] = tmp[i]
	}
	return string(tmp)
}

func smallerPalindromic(s string) string {
	if isParlindromic(s) {
		return s
	}

	tmp := []byte(s)
	for i := 0; i < len(tmp)/2; i++ {
		tmp[len(tmp)-1-i] = tmp[i]
	}
	t := string(tmp)
	sI, _ := strconv.Atoi(s)
	tI, _ := strconv.Atoi(t)
	if sI > tI {
		return t
	}
	for i := (len(tmp) - 1) / 2; i >= 0; i-- {
		if tmp[i] != '0' {
			tmp[i] -= 1
			break
		} else {
			tmp[i] = '9'
		}
	}
	if tmp[0] == '0' {
		for i := 0; i < len(tmp); i++ {
			tmp[i] = '9'
		}
		return string(tmp[1:])
	}
	for i := 0; i < len(tmp)/2; i++ {
		tmp[len(tmp)-1-i] = tmp[i]
	}

	return string(tmp)
}

func isParlindromic(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func larger(n string) string {
	t, _ := strconv.Atoi(n)
	return strconv.Itoa(t + 1)
}

func smaller(n string) string {
	t, _ := strconv.Atoi(n)
	return strconv.Itoa(t - 1)
}

func nearestPalindromic(n string) string {
	larger := largerPalindromic(larger(n))
	smaller := smallerPalindromic(smaller(n))
	fmt.Println("larger", larger)
	fmt.Println("smaller", smaller)
	largerValue, _ := strconv.Atoi(larger)
	smallerValue, _ := strconv.Atoi(smaller)
	thisValue, _ := strconv.Atoi(n)
	if largerValue-thisValue >= thisValue-smallerValue {
		return smaller
	} else {
		return larger
	}
}
func main() {
	fmt.Println(nearestPalindromic("100"))
}
