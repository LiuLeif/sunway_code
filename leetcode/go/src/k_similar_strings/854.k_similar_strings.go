// 2018-12-10 13:41
package main

import "fmt"

type Tuple struct {
	x string
	y string
	z int
}

type Queue []Tuple

func (this *Queue) Poll() Tuple {
	tmp := (*this)[0]
	*this = (*this)[1:]
	return tmp
}

func (this *Queue) Offer(x Tuple) {
	*this = append(*this, x)
}

func (this *Queue) Empty() bool {
	return len(*this) == 0
}

func (this *Queue) Peek() Tuple {
	return (*this)[0]
}

func kSimilarity(A string, B string) int {
	queue := Queue{}
	queue.Offer(Tuple{A, B, 0})
	for !queue.Empty() {
		top := queue.Poll()
		a, b := top.x, top.y
		if a == b {
			return top.z
		}
		// abc,acb
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				a, b = a[i:], b[i:]
				break
			}
		}
		// bc,cb
		for i := 0; i < len(a); i++ {
			if b[i] == a[0] {
				bb := []byte(b)
				bb[i], bb[0] = bb[0], bb[i]
				queue.Offer(Tuple{a[1:], string(bb[1:]), top.z + 1})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(kSimilarity("abc", "bca"))
	fmt.Println(kSimilarity("ab", "ba"))
}
