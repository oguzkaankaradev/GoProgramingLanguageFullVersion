package bank

import "fmt"

func Bank() {
	fmt.Println("Bankamıza hoşgeldiniz")
	fmt.Println("Ne yapmak istersiniz")
	fmt.Println("1. Bakiye kontrol")
	fmt.Println("2. Depositler")
	fmt.Println("3. Para çekme")
	fmt.Println("4. kartı al")

	var choice int
	fmt.Println("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {

	}

	fmt.Println("Your choice:", choice)
}
