// 2018-10-23 10:19
package main

import "fmt"
import "math/rand"
import "math"

type Solution struct {
	radius  float64
	xCenter float64
	yCenter float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {
	deg := rand.Float64() * 2 * math.Pi
	length := math.Sqrt(rand.Float64()) * this.radius

	return []float64{length*math.Cos(deg) + this.xCenter, length*math.Sin(deg) + this.yCenter}

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */

func main() {
	obj := Constructor(10, 5, -7.5)
	fmt.Println(obj.RandPoint())
	fmt.Println(obj.RandPoint())
	fmt.Println(obj.RandPoint())
	fmt.Println(obj.RandPoint())
	fmt.Println(obj.RandPoint())
}
