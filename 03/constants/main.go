package main

import (
	"fmt"
	fmt2 "fmt"

	// import aliasName "package_to_import"

	"os"
)

const (
	x      = 10
	y      = 20
	pi int = 300
)

func main() {

	const x = 50
	fmt2.Println(x)

	// x = 20
	// Cannot do this, because x is constant

	fmt.Println()

	// var name = "Bala"
	name := "Bala"
	fmt.Println("Hi this is ", name)

	os.Exit(-1)
}
