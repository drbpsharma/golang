package main

import (
	"fmt"
)
const(
	_=iota
	KB=1 << (10*iota)
	MB
	GB
)

func main() {
	fmt.Printf("%v %v %v\n",KB,MB,GB) 
}
