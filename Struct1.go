package main

import (
	"fmt"
)
type Book struct{
	bookno int
	title string
	author string
}
func main() {
	b := Book{bookno:101,title:"Go Programming",author:"Dr B P Sharma"}
	fmt.Println(b) //{101 Go Programming Dr B P Sharma}
}
