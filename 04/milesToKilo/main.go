package main

import (
	"fmt"
	"os"
)

//  "MilesToKilometers ... "
const MilesToKilometers = 1.60934

func main() {

	var miles float64

	fmt.Printf("Enter miles ")

	_, err := fmt.Scanf("%v", &miles)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	kilometers := miles * MilesToKilometers

	fmt.Printf("Kilometers = %f \n", kilometers)

	fmt.Println()
}
