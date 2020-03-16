package main

import (
	"fmt"
)

func main() {
	var n int=6
	var s int
	var c int
	multiresult(n,&s,&c) //passing address
	fmt.Printf("Square is %v, Cube is %v",s,c)
}

func multiresult(n int, square *int, cube *int){
	*square=n*n
	*cube=n*n*n
}
