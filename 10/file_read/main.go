package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File reading")

	var filename = "./main.go"

	/*
		dat, err := ioutil.ReadFile(filename)

		if err != nil {
			panic(err)
		}

		fmt.Print(string(dat))
	*/

	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	b1 := make([]byte, 1000)
	n1, err := f.Read(b1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d bytes: %s\n", n1, string(b1))

}
