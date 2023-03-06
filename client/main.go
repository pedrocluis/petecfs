package main

import (
	"fmt"
	"time"
)

func registerUser(username string, password string) int {
	return 1
}

func main() {
	fmt.Println("Welcome to petecfs!")
	time.Sleep(2 * time.Second)
	cleanTerminal() // Clears the screen

	//Ask for login or register a new account
	userId := registerLoginMenu()

	//User exits app
	if userId == 0 {
		return
	}

	fmt.Printf("%d\n", userId)
}
