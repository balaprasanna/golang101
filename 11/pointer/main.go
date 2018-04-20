package main

import "fmt"

func changeMe(name *string) {
	*name = "Prasanna"
}

func main() {

	fmt.Println("Pointer example")

	var name = "bala"
	var ptr = &name

	changeMe(ptr)
	fmt.Println("Name = ", name)
}
