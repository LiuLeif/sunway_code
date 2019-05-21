// 2018-12-03 13:12
package main

import "fmt"

func largestTimeFromDigits(A []int) string {
	// 1 2 3 4
	hour := 0
	minute := 0
	maxHour, maxMinute := -1, -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			hour = A[i]*10 + A[j]
			if hour >= 24 {
				continue
			}
			for m := 0; m < 4; m++ {
				for n := 0; n < 4; n++ {
					if m == n || m == i || m == j || n == i || n == j {
						continue
					}
					minute = A[m]*10 + A[n]
					if minute >= 60 {
						continue
					}
					if maxHour*60+maxMinute < hour*60+minute {
						maxHour = hour
						maxMinute = minute
					}
				}
			}

		}
	}

	if maxHour == -1 {
		return ""
	}
	return fmt.Sprintf("%02d:%02d", maxHour, maxMinute)
}

func main() {
	fmt.Println(largestTimeFromDigits([]int{1, 2, 3, 2}))
	fmt.Println(largestTimeFromDigits([]int{0, 2, 3, 9}))
}
