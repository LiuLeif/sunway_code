// 2018-01-29 10:18
package main

import (
	"fmt"
	"strconv"
)

func solveEquation(equation string) string {
	equation += string('+')
	negative := false
	digit := 0
	hasX := false
	hasDigit := false
	totalCoefficient := 0
	totalDigit := 0
	reverse := false

	calc := func() {
		if !hasDigit && hasX {
			digit = 1
		}
		if reverse {
			negative = !negative
		}
		if negative {
			digit = -digit
		}
		if hasX {
			totalCoefficient += digit
		} else {
			totalDigit += digit
		}
		digit = 0
		hasX = false
		hasDigit = false
		negative = false
	}

	for i := 0; i < len(equation); i++ {
		c := equation[i]
		switch c {
		case '+':
			calc()
		case '-':
			calc()
			negative = true
		case '=':
			calc()
			reverse = true
		case 'x':
			hasX = true
		default:
			hasDigit = true
			digit = digit*10 + int(c-'0')
		}
	}
	if totalCoefficient == 0 {
		if totalDigit == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	} else {
		return "x=" + strconv.Itoa(-totalDigit/totalCoefficient)
	}
}
func main() {
	fmt.Println(solveEquation("0x=0"))
}
