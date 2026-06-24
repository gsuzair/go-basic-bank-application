package main

import "fmt"

func main() {
	var accountBalance = 1000.00

	fmt.Println("Welcome to go Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit amount: ")
		var depositAmount float64
		fmt.Scanln(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Deposit amount must be greater than zero.")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Your new balance is: ", accountBalance)
	} else if choice == 3 {
		fmt.Print("Your withdrawn amount: ")
		var withdrawAmount float64
		fmt.Scanln(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Withdraw amount must be greater than zero.")
			return
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds. Your balance is: ", accountBalance)
			return
		}

		accountBalance -= withdrawAmount
		fmt.Println("Your new balance is: ", accountBalance)
	} else {
		fmt.Println("Thank you for using go Bank. Goodbye!")
	}

	fmt.Println("Your choice: ", choice)

}
