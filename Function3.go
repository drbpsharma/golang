package main

import (
	"fmt"
	"errors"
)

func main() {
	n:=-6
	f,e:=factorial(n)
	if(e==nil){
		fmt.Printf("Factorial is %v",f)
	}else{
		fmt.Printf("Error : %v",e)
	}
}
func factorial(n int) (int,error){
	if n<0{
		return 0,errors.New("Invalid Number. Number must be >=0")
	}
	f:=1
	i:=1
	for i<=n{
		f=f*i
		i++
	}
	return f,nil
}
