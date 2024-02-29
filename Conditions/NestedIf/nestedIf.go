package nestedIf

import (
	"fmt"
)

func ATM() {
	var accountBalance float64 = 1000.0
	fmt.Println("Bankamıza hoşgeldiniz")
	fmt.Println("Ne yapmak istersiniz")
	fmt.Println("1. Bakiye kontrol")
	fmt.Println("2. Para yatırma")
	fmt.Println("3. Para çekme")
	fmt.Println("4. kartı al")

	var choice int
	fmt.Print("Your choice : ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		fmt.Print("Para tutarı :")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("invalid amount. Must be greater than 0.")
			return
		}

		accountBalance += float64(depositAmount)
		fmt.Println("Balance updated! :", accountBalance)

	} else if choice == 3 {
		fmt.Println("withdrawal amount : ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)

		if withdrawalAmount <= 0 {
			fmt.Println("invalid Maount. Must be greater than 0.")
			return
		}

		if withdrawalAmount > float64(accountBalance) {
			fmt.Println("Invalid amount. You cant withdraw more than you have.")
			return

		}

		accountBalance -= float64(withdrawalAmount)
		fmt.Println("Balance Updated! new amount : ", accountBalance)
	} else {
		fmt.Println("See you")
	}
}
