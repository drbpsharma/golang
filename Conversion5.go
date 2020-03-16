package main

import (
	"fmt"
)

func main() {
	n:=complex(3,5)
	fmt.Printf("%v %T\n",n,n)
	fmt.Printf("Real Part %v, Imagenary Part %v",real(n), imag(n))
	
}
