package main

import (
	"fmt"
)

// var email string = "likhitviwatkul_t@su.ac.th"
func main() {
	//var name string = "thammatorn"
	var age int = 20
	email := "likhitviwatkul_t@su.ac.th"
	gpa := 2.86
	firstname, lastname := "thammatorn", "likhitviwatkul"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n", firstname, lastname, age, email, gpa)
}
