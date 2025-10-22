package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello world")
	var input string

	fmt.Println("What your name")
	fmt.Scan(&input)
	fmt.Printf("Hello, %s!\n", input)
}

//
