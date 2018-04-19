package main

import "fmt"

func main() {
	fmt.Println("If Else example")

	i := 0
	fmt.Print("Enter your age: ")
	fmt.Scanf("%v", &i)
	fmt.Println()
	if i < 18 {
		fmt.Println("Young")
	} else if i < 30 {
		fmt.Println("Matured")
	} else {
		fmt.Println("Old")
	}
}
