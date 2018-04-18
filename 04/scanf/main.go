package main

import (
	"fmt"
)

func main() {

	var name string

	fmt.Printf("Enter your name ")

	fmt.Scanf("%s", &name)

	fmt.Println("Welcome Mr. ", name)

	var x = 10
	var y = 20
	fmt.Printf("Result x * y = %d \n", x*y)
}
