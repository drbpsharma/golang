package main

import (
	"fmt"
)

func main() {
	var a int8=4
	var b int=5
	fmt.Println(a+b) // Error - ./prog.go:10:15: invalid operation: a + b (mismatched types int8 and int)
}
