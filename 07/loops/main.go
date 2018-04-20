package main

import "fmt"
import "reflect"

func main() {
	fmt.Println("Looping example: ")
	fmt.Println("Type 1")

	count := 20

	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	fmt.Println("Type 2:")
	i := 0

	var x = 10.34
	// var y int = 20
	fmt.Println("Assert x=y ", reflect.TypeOf(x).Kind() == reflect.Int)

	for i < count {
		fmt.Printf("%v", i)
		i++
	}

	fmt.Println()

	fmt.Println("Type 3:")

	for {
		fmt.Println(i)
		break
	}

	fmt.Println("Type 4:")

	msg := "hello"
	for i, value := range msg {
		fmt.Println(i, value)

	}

	fmt.Println("Type 5:")

	cond := true

	for cond {
		fmt.Println("Hi")
		break
	}
}
