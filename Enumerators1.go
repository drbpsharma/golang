package main

import (
	"fmt"
)
const(
	zero=iota
	one
	two
	three
)

const(
	hr=iota+101
	marketing
	finance
	production
)
func main() {
	fmt.Printf("%v %v %v %v\n",zero,one,two,three) // 0 1 2 3
	fmt.Printf("%v %v %v %v\n",hr,marketing,finance,production) // 101 102 103 104
}
