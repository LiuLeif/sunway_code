// 2018-11-01 18:00
package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPalindromic(N int) bool {
	s := strconv.Itoa(N)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func largerPalindromic(N int) int {
	if isPalindromic(N) {
		return N
	}
	s := strconv.Itoa(N)
	tmp := []byte(s)
	for i := 0; i < len(tmp)/2; i++ {
		tmp[len(tmp)-1-i] = tmp[i]
	}
	t := string(tmp)
	sValue, _ := strconv.Atoi(s)
	tValue, _ := strconv.Atoi(t)
	if sValue < tValue {
		ret, _ := strconv.Atoi(t)
		return ret
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
	ret, _ := strconv.Atoi(string(tmp))
	return ret
}

func isPrime(N int) bool {
	if N == 1 {
		return false
	}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(N)))); i++ {
		if N%i == 0 {
			return false
		}
	}
	return true
}

func primePalindrome(N int) int {
	if isPrime(N) && isPalindromic(N) {
		return N
	}
	for {
		N = largerPalindromic(N + 1)
		if isPrime(N) {
			return N
		}
	}
	return 0
}

func main() {
	fmt.Println(primePalindrome(1))
}
