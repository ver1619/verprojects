package main

import "fmt"

const username = "admin"
const password = "admin123"
const pin = "4321"
const accountNumber = "1234567890"

const amount = 500000

func login(u, p string) bool {
	if u == username && p == password {
		return true
	}
	return false
}

func atm() {
	fmt.Println("Welcome to Simple ATM!", username)
}

func accountDetails() {
	fmt.Println("Fetching Account Details...")
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
	fmt.Println("PIN:", pin)
	fmt.Println("Account Number:", accountNumber)
}

func withdraw() {
	var withdrawAmount int
	fmt.Print("Enter amount to withdraw: ")
	fmt.Scanln(&withdrawAmount)

	if withdrawAmount > amount {
		fmt.Println("Insufficient balance!")
	} else {
		fmt.Println("Please collect your cash:", withdrawAmount)
	}
	fmt.Println("Amount withdrawn successfully!")
	fmt.Println("Updating balance...")
	fmt.Printf("Remaining balance: %d\n", amount-withdrawAmount)
}

func deposit() {
	var depositAmount int
	fmt.Print("Enter amount to deposit: ")
	fmt.Scanln(&depositAmount)
	fmt.Println("Amount to be deposited:", depositAmount)
	fmt.Println("Amount deposited successfully!")
	fmt.Println("Updating balance...")
	fmt.Printf("New balance: %d\n", amount+depositAmount)
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
		atm()

		for {
			fmt.Print("\nDashboard:\n")
			fmt.Println("1. Account Details")
			fmt.Println("2. Withdraw Amount")
			fmt.Println("3. Deposit Amount")
			fmt.Println("4. Exit")
			var choice int
			fmt.Print("Select an option : ")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				accountDetails()
			case 2:
				withdraw()
			case 3:
				deposit()
			case 4:
				fmt.Println("Exiting... Thank you for using Simple ATM!")
				return
			default:
				fmt.Println("Invalid option. Please try again.")
			}
		}
	} else {
		fmt.Println("Login failed!")
		fmt.Println("FORGOT PASSWORD?")
		fmt.Println("Type 'yes' to retrieve your credentials or 'no' to exit:")
		var choice string
		fmt.Scanln(&choice)

		if choice == "yes" {
			fmt.Println("Provide Valid PIN to fetch details:")
			var inputPin string
			fmt.Scanln(&inputPin)

			if inputPin != pin {
				fmt.Println("Invalid PIN. Exiting...")
				return
			} else {
				fmt.Println("Valid PIN. Fetching details...")
			}
			fmt.Println("Fetching user credentials...")
			fmt.Println("Username:", username)
			fmt.Println("Password:", password)
			fmt.Println("Go back to Login")
			var goBack string
			fmt.Scanln(&goBack)
			if goBack == "yes" {
				main()
			} else {
				fmt.Println("Exiting...")
			}
		} else {
			fmt.Println("Exiting...")
		}
	}
}
