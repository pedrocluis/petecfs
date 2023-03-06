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
	menu.AddItem("Exit", "exit")
	choice := menu.Display()
	cleanTerminal()
	return choice
}

func yesNoMenu(prompt string) string {
	menu := gocliselect.NewMenu(prompt)
	menu.AddItem("Yes", "yes")
	menu.AddItem("No", "no")
	choice := menu.Display()
	cleanTerminal()
	return choice
}

// Returns -1 if log in failed or user id if log in successful
func checkLogin(user string, password string) int {
	return 1
}

// Returns user id or -1 if log in failed
func login() int {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	user, _ := reader.ReadString('\n')
	fmt.Print("Password: ")
	password, _ := term.ReadPassword(0)
	cleanTerminal()
	return checkLogin(user, string(password))
}

// Returns -1 if invalid username, 0 if username taken, 1 if available username
func checkUsername(username string) int {
	return 1
}

func checkPassword(username string) bool {
	return true
}

func register() int {

	user := ""
	u_id := 0

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Username: ")
		user, _ = reader.ReadString('\n')
		flagUser := checkUsername(user)
		if flagUser == -1 {
			cleanTerminal()
			retry := yesNoMenu("Invalid Username, try again?")
			if retry == "no" {
				return -1
			}
		}
		if flagUser == 0 {
			cleanTerminal()
			retry := yesNoMenu("Username taken, try again?")
			if retry == "no" {
				return -1
			} else {
				continue
			}
		}
		if flagUser == 1 {
			break
		}
	}

	for {
		fmt.Println("Password must contain at least 6 characters with minimum one symbol or number")
		fmt.Print("Password: ")
		password, _ := term.ReadPassword(0)
		fmt.Print("\nConfirm password: ")
		dupPassword, _ := term.ReadPassword(0)
		if string(password) != string(dupPassword) {
			cleanTerminal()
			retry := yesNoMenu("Passwords don't match, try again?")
			if retry == "no" {
				return -1
			} else {
				continue
			}
		}
		if !checkPassword(string(password)) {
			cleanTerminal()
			retry := yesNoMenu("Invalid password, try again?")
			if retry == "no" {
				return -1
			} else {
				continue
			}
		}
		u_id = registerUser(user, string(password))
		fmt.Println("")
		fmt.Println("Register successful!")
		time.Sleep(1 * time.Second)
		cleanTerminal()
		break
	}

	return u_id

}

// Returns 0 if user wants to exit app or user_id
func registerLoginMenu() int {
	userId := -1
	for {
		option := loginMenu()

		//Try to login
		if option == "login" {
			for {
				userId = login()
				//Login failed
				if userId == -1 {
					loginRetryYesNo := yesNoMenu("Log in failed, try again?")
					if loginRetryYesNo == "no" {
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
		} else if option == "register" {
			userId = register()
			break
		} else if option == "exit" {
			userId = 0
			break
		}
	}
	return userId
}
