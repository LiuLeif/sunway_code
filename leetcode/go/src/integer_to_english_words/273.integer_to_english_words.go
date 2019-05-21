// 2018-12-14 15:11
package main

import (
	"fmt"
	"strings"
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	lessThanTwenty := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands := []string{"", "Thousand", "Million", "Billion"}

	var helper func(x int) string
	helper = func(x int) string {
		// 123 -> one hundred twenty three
		ret := ""
		if x < 20 {
			ret = lessThanTwenty[x] + " "
		} else if x < 100 {
			ret = tens[x/10] + " " + lessThanTwenty[x%10] + " "
		} else {
			ret = lessThanTwenty[x/100] + " Hundred " + helper(x%100)
		}
		return strings.Trim(ret, " ")
	}
	// 1234567
	ret := ""
	step := 0
	for num > 0 {
		if num%1000 != 0 {
			ret = helper(num%1000) + " " + thousands[step] + " " + ret
		}
		step++
		num /= 1000
	}

	return strings.Trim(ret, " ")
}

func main() {
	// 1 234 567
	// One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven
	fmt.Println(numberToWords(50868))
	fmt.Println(numberToWords(123456))
	fmt.Println(numberToWords(1234))
}
