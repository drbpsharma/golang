package main

import (
	"fmt"
)

func main() {
	var a int=10
	var b int=3
	fmt.Println(a==b) // false
	fmt.Println(a!=b) // true
	fmt.Println(a>b) // true
	fmt.Println(a>=b) // true
	fmt.Println(a<b) // false
	fmt.Println(a<=b) // false
}
