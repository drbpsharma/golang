package main

import (
	"fmt"
)

func main() {
	var y int = 2020
	var result bool=y%4==0 && y%100!=0 || y%400==0
	fmt.Println(result) //Prints true
}
