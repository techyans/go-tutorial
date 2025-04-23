package main

import "fmt"

func main() {
	var accountBalance = 20000.0
	fmt.Println("Welcome to Go Bank")
	fmt.Println("Select your choice")
	fmt.Println("1. Deposit Money")
	fmt.Println("2. Withdraw Money")
	fmt.Println("3. Check Balance")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your Choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var depamount float64
		fmt.Print("Enter the amount you want to deposit:")
		fmt.Scan(&depamount)
		if depamount <= 0 {
			fmt.Println("Invalid amount")
			return
		}
		fmt.Println("Your deposit amount is:", depamount)
		fmt.Println("Your new balance is:", accountBalance+depamount)
	case 2:
		var withamount float64
		fmt.Print("Enter the amount you want to withdraw:")
		fmt.Scan(&withamount)
		if withamount > accountBalance {
			fmt.Println("Insufficient balance")
			return
		} else if withamount <= 0 {
			fmt.Println("Invalid amount")
			return
		}

		fmt.Println("Your withdraw amount is:", withamount)
		fmt.Println("Your new balance is:", accountBalance-withamount)
	case 3:
		fmt.Printf("Your balance is: %.2f\n", accountBalance)
	default:
		fmt.Println("Goodbye")
		fmt.Println("Thank you for using Go Bank")
	}
	/*for i := 0; i < 200; i++ {
		fmt.Println("Select your choice")
		fmt.Println("1. Deposit Money")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 3 {
			fmt.Printf("Your balance is: %.2f\n", accountBalance)
		} else if choice == 1 {
			var depamount float64
			fmt.Print("Enter the amount you want to deposit:")
			fmt.Scan(&depamount)
			if depamount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			fmt.Println("Your deposit amount is:", depamount)
			fmt.Println("Your new balance is:", accountBalance+depamount)
			break
		} else if choice == 2 {
			var withamount float64
			fmt.Print("Enter the amount you want to withdraw:")
			fmt.Scan(&withamount)
			if withamount > accountBalance {
				fmt.Println("Insufficient balance")
				continue
			} else if withamount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			fmt.Println("Your withdraw amount is:", withamount)
			fmt.Println("Your new balance is:", accountBalance-withamount)
			break
		} else {
			fmt.Println("Goodbye")
			break
		}
	}

	fmt.Println("Thank you for using Go Bank")*/
}
