package main

import (
	"fmt"
)

func main() {
	//Map ve arraylerimizin üzeridne daha kolay gezinmemizi sağlamaktadır.

	// var numbers = []int{143, 2235, 313, 43454}

	// for index := 0; index < len(numbers); index++ {
	// 	fmt.Println(numbers[index])
	// }

	// FOREACH İLE YAZALIM
	//index = array veya slicenin indexi örbeğin sıfırıncı index [0]
	//value = indeximizin valuesi
	// range = neyin üzerinde gezeceğiz, hangi array slice veya map veya list.
	// kullanmadığımız değerler olduğunda _ kullanmamız bestpractis açısından iyidir.

	// for index, value := range numbers {
	// 	fmt.Println(index, value)
	// }

	// for index, _ := range numbers {
	// 	fmt.Println(index)
	// }

	// for _, value := range numbers {
	// 	fmt.Println(value)
	// }

	var name = "muhammet oraz"

	for _, karakter := range name {
		//fmt.Println(karakter)
		fmt.Printf("%c\n", karakter)
	}

}
