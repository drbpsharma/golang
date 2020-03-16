package main

import (
	"fmt"
)

func main() {
	var a int8=5
	var b int16=6
	var c int32=7
	var d int64=8
	fmt.Printf("%v %v %v %v\n",a,b,c,d) //5 6 7 8
	
	var e uint8=5
	var f uint16=6
	var g uint32=7
	var h uint64=8
	fmt.Printf("%v %v %v %v\n",e,f,g,h) //5 6 7 8
	
	var i float32=3.4
	var j float64=5.6
	fmt.Printf("%v %v\n",i,j) //(3+5i) (3+5i)
	
	var k complex64=3+5i
	var l complex128=3+5i
	fmt.Printf("%v %v\n",k,l) //(3+5i) (3+5i)
	
	var m string="Hello";
	fmt.Printf("%v\n",m) //Hello
	
	var n bool=true
	fmt.Printf("%v\n",n) //true
	
	var o rune='A'
	fmt.Printf("%v %T\n",o,o) //65 int32
	fmt.Printf("%v %T",string(o),string(o)) //A string
}
