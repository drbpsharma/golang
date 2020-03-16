package main

import (
	"fmt"
)

func main() {
	var a int
	var b float64
	var c bool
	var d rune
	var e string
	fmt.Printf("%v %T\n",a,a) //0 int
	fmt.Printf("%v %T\n",b,b) //0 float64
	fmt.Printf("%v %T\n",c,c) //false bool
	fmt.Printf("%v %T\n",d,d) //0 int32
	fmt.Printf("%v %T\n",e,e) // string
}
