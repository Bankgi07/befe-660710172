package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password123"

	// Hash password (Cost Factor = 12)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err)
	}

	fmt.Println("Password:", password)
	fmt.Println("Hash:", string(hash))

	// Verify password
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == nil {
		fmt.Println("Password correct!")
	} else {
		fmt.Println("Password wrong!")
	}

	// ลองรหัสผ่านผิด
	err = bcrypt.CompareHashAndPassword(hash, []byte("wrongpassword"))
	if err == nil {
		fmt.Println("Password correct!")
	} else {
		fmt.Println("Password wrong!")
	}

	password = "password123"
	hash, err = bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nPassword:", password)
	fmt.Println("Hash:", string(hash))

	// ===========================================================

	password = "admin123"
	hash, err = bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nPassword:", password)
	fmt.Println("Hash:", string(hash))

	password = "editor123"
	hash, err = bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nPassword:", password)
	fmt.Println("Hash:", string(hash))

	password = "user123"
	hash, err = bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nPassword:", password)
	fmt.Println("Hash:", string(hash))
}
