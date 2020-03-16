package main

import (
	"fmt"
)

func main() {
	n:=2+5i
	fmt.Printf("%v %T\n",n,n)
	fmt.Printf("Real Part %v, Imagenary Part %v",real(n), imag(n)) //extracting real and imagenary parts from a complex number
	
}
