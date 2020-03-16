package main

import (
	"fmt"
)

func main() {
	s:="Dr B P Sharma"
	fmt.Printf("%v %T\n",s,s)	//Dr B P Sharma string
	fmt.Printf("%v %T\n",s[1],s[1])	//114 uint8
	b:=[]byte(s)
	fmt.Printf("%v %T\n",b,b)	//[68 114 32 66 32 80 32 83 104 97 114 109 97] []uint8
}
