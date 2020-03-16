package main

import (
	"fmt"
)

func main() {
	n:=6
	f:=factorial(n)
	fmt.Printf("Factorial is %v",f)
}
func factorial(n int) int{
	f:=1
	i:=1
	for i<=n{
		f=f*i
		i++
	}
	return f
}
