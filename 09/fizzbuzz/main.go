// Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)
package main

import "fmt"

func main() {
	fmt.Println("FizzBuzz")

	for num := 1; num < 100; num++ {
		msg := ""
		if num%3 == 0 {
			msg = msg + "Fizz"
		}
		if num%5 == 0 {
			msg = msg + "Buzz"
		}
		if msg != "" {
			fmt.Println(msg, num)
		}

	}
}
