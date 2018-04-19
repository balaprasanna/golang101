package main

import "fmt"
import "os"

func main() {
	fmt.Println(
		`Please select your prefered choice: 
	1. Miles to Kilometers
	2. Fahrenheit to Celsius
	3. Pounds to Kilograms
`)

	fmt.Printf("Your Option = ")
	var choice int64

	_, err := fmt.Scanf("%d", &choice)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if choice < 1 || choice > 3 {
		fmt.Println("Incorrect Option:")
		os.Exit(-1)
	}

	switch choice {
	case 1:
		fmt.Printf("Choice = %v \n", choice)
		miles := collectInputs("Miles")
		kilometers := convert(miles, 1.60934)
		fmt.Printf("%v Miles = %v Kilometers \n", miles, kilometers)
	case 2:
		fmt.Printf("Choice = %v \n", choice)
		farenheit := collectInputs("Fahrenheit")
		clecius := convert(farenheit-32, 5.0/9.0)
		fmt.Printf("%v Farenheit = %v Clecius \n", farenheit, clecius)
	case 3:
		fmt.Printf("Choice = %v \n", choice)
		pounds := collectInputs("Pounds")
		kilograms := convert(pounds, 0.453592)
		fmt.Printf("%v Pounds = %v Kilograms \n", pounds, kilograms)
	}

}

func collectInputs(Type string) float64 {
	var input float64
	fmt.Printf("Please enter values in %v : ", Type)
	fmt.Scanf("%f", &input)
	return input
}

func convert(input float64, scale float64) float64 {
	return (input * scale)
}
