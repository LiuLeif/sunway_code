// 2018-01-24 13:42
package main

import "fmt"

func reverseBit(n uint32) uint32 {
	ret := uint32(0)
	for i := uint32(0); i < 32; i++ {
		ret = ret<<1 | (n>>i)&0x1
	}
	return ret
}

func main() {
	fmt.Println(reverseBit(43261596))
}
