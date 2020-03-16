package main

import (
	"fmt"
)

func main() {
	a:=5
	b:=5.0
	c:='A'
	d:="A"
	fmt.Printf("%v %T\n",a,a) //5 int
	fmt.Printf("%v %T\n",b,b) //5 float64
	fmt.Printf("%v %T\n",c,c) //65 int32
	fmt.Printf("%v %T\n",d,d) //A string

}
