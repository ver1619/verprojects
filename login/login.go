package main

import "fmt"

const username = "admin"
const password = "admin123"

func login(u, p string) bool {
	if u == username && p == password {
		return true
	}
	return false
}

func main() {
	var user string
	var pass string

	fmt.Print("Enter username: ")
	fmt.Scanln(&user)
	fmt.Print("Enter password: ")
	fmt.Scanln(&pass)

	if login(user, pass) {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed!")
		fmt.Println("Want to fetch user credentials? (yes/no)")
		var choice string
		fmt.Scanln(&choice)

		if choice == "yes" {
			fmt.Println("Fetching user credentials...")
			fmt.Println("Username:", username)
			fmt.Println("Password:", password)
		} else {
			fmt.Println("Exiting...")
		}
	}
}
