package main

import "fmt"

// This is invalid
// var x

// This is valid
// var x string
// var x = ""

var PI = 3.14

var msg string = "Welcome everyone"

func main() {
	// Both of them is same
	// x := "HI"
	// var x = "HI"

	x := 10
	fmt.Println(x)

	// This is invalid, Because we cannot create a same variable name again.string
	// x := 70

	// Correct way to assign is
	x = 70

	fmt.Println(msg)

	scope()
}

func scope() {

	// This x is local to this function block
	x := 200

	fmt.Println(x)

	{
		y := 10
		fmt.Println(y)
	}
	// fmt.Println(y)
	// This is line in invalid,
	// Because y is not available as function scope.
	// Y is part of a block scope , enclosed within the {} braces

	fun()
}
