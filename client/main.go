package main

import (
	"bufio"
	"fmt"
	"github.com/nexidian/gocliselect"
	"golang.org/x/term"
	"os"
	"time"
)

func cleanTerminal() {
	fmt.Print("\033[H\033[2J")
}

func loginMenu() string {
	menu := gocliselect.NewMenu("Sign in or register")
	menu.AddItem("Sign In", "login")
	menu.AddItem("Register", "register")
	choice := menu.Display()
	cleanTerminal()
	return choice
}

func loginYesNoMenu() string {
	menu := gocliselect.NewMenu("Log in failed, try again?")
	menu.AddItem("Yes", "yes")
	menu.AddItem("No", "no")
	choice := menu.Display()
	cleanTerminal()
	return choice
}

func checkLogin(user string, password string) int {
	return -1
}

// Returns user id or -1 if log in failed
func login() int {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	user, _ := reader.ReadString('\n')
	cleanTerminal()
	fmt.Print("Password: ")
	password, _ := term.ReadPassword(0)
	cleanTerminal()
	return checkLogin(user, string(password))
}

func registerLoginMenu() int {
	userId := -1
	for {
		option := loginMenu()

		//Try to login
		if option == "login" {
			for {
				userId = login()
				if userId == -1 {
					yes_no := loginYesNoMenu()
					if yes_no == "no" {
						break
					} else {
						continue
					}
				} else {
					break
				}
			}
			if userId == -1 {
				continue
			} else {
				break
			}
		}
	}
	return userId
}

func main() {
	fmt.Println("Welcome to petecfs!")
	time.Sleep(3)
	cleanTerminal() // Clears the screen

	//Ask for login or register a new account
	userId := registerLoginMenu()

	fmt.Printf("%d\n", userId)
}
