package main

import (
	"fmt"
)

func main() {
	var a int=10		// 1010
	var b int=3		// 0011
	fmt.Println(a&b) 	// 0010 - 2
	fmt.Println(a|b) 	// 1011 - 11
	fmt.Println(a^b) 	// 1001 - 9
	fmt.Println(a&^b) 	// 0100 - 8
	fmt.Println(5 << 3)	// 101 << 3 = 101000 = 40
	fmt.Println(50 >> 3)	// 110010 >>3 = 000110 = 6
}
