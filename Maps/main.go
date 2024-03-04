package main

import "fmt"

// java c görebilirsiniz. Ayrıca python, c# sözlükler dede görebilirsiniz.
// map : verilerimizi key value şeklinde tutmaktadır. bu bir veritipidir.
// key : user.name / value: muhammet - oraz - ahmet
//sözlük = key value
// mapplerde sıralama farklı olabilir. çünkü bellekte mapler farklı tutulabilir.

func main() {
	// her nameye karşılık gelen student number olsun
	var names map[string]int

	names = make(map[string]int, 0)
	names["oraz"] = 1206
	names["muhammet"] = 5012

	fmt.Println(names)

	// key = isim value = soyisim

	var fullNames map[string]string

	fullNames = make(map[string]string, 0) // map i yazdırabilmek için gerekli

	fullNames["oraz"] = "berdimuradov"
	fullNames["a"] = "z"
	fullNames["muhammet"] = "garabekov"

	//fmt.Println(fullNames)
	//fmt.Println(fullNames["a"])

	// olmayan değer çağıralım
	fmt.Println(fullNames["b"])

	// key name value price
	productName := map[string]int{
		"saat":       23,
		"bilgisayar": 12323,
		"telefon":    123123, // en sonundada virgül olmak zorunda.
	}

	fmt.Println(productName)

}
