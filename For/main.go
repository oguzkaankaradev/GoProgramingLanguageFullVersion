package main

import (
	"fmt"
	atm "oneVideoModule/For/BankApp"
)

func main() {
	// index := 1

	// for index <= 10 {
	// 	fmt.Println(index)

	// 	//index = index + 1
	// 	//index++
	// }

	//ikinci yöntem
	// for index := 1; index <= 10; index++ {
	// 	fmt.Println(index)
	// }

	//example toplamlarını yazdır.

	// total := 0
	// for index := 1; index <= 10; index++ {
	// 	//total = total + index
	// 	total += index
	// }
	// fmt.Println(total)

	//example 2 : array slice
	//len = toplama işlemini yapıyor.
	index := 0
	var users = []string{"ali", "ahmet", "tgffg", "tyh", "uku", "oıl", "dfgtry", "ahmfet", "ytj", "ahmet", "ghjuj", "fghy", "getg"}
	for index < len(users) {
		fmt.Println(users[index])
		index++
	}

	atm.AtmCode()

}
