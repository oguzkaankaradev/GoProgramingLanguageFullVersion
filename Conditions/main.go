package main

import "fmt"

func main() {
	var age int = 20

	if age >= 18 {
		fmt.Println("Bu uçağa binebilirsiniz biletin yoksa lütfen bilet alınız")
	} else if age >= 16 && age < 18 {
		fmt.Println("Uçağa binebilirsin fakat bilet alamazsın")
	} else {
		fmt.Println("Büyüde gel")
	}
}
