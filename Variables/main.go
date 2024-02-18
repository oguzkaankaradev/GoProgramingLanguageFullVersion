package main

import (
	"fmt"
	"reflect"
)

func main() {

	//first method
	var productName string
	var allowAge int
	var discount float32 //float64
	var isActive bool    //true veya false

	productName = "goLang"
	allowAge = 15
	discount = 1.34
	isActive = false

	//Education Detail : Tek bir print içerisine birden fazla değişken yazdırabilirsin.
	//fmt.Println(productName, " ", allowAge, " ", discount, " ", isActive)

	//

	fmt.Println("Eğitimin ismi :", productName, reflect.TypeOf(productName))
	fmt.Println("Hangi yaştaki öğrenciler girecek : ", allowAge, reflect.TypeOf(allowAge))
	fmt.Println("indirim : ", discount, reflect.TypeOf(discount))
	if isActive == true {
		fmt.Println("Eyy oraz kursa kayıt olabilirsin")
	} else {
		fmt.Println("Lütfen kursun açılmasını takipediniz")
	}
	//reflect.TypeOf(productName) : typeof veri tipinin döndürür.

	//second  method

	var deathVariables string = "Değişkenin ömrü bu sayfada global değil bu."

	println(deathVariables)

	//third method

	var veritipiniKendiBelirleyenYontem = "Senin veri tipin string olsun. Dikkat et veri tipinin yanına bişey yazmadık.eğer buraya sayı yazsaydık int oalcaktı ama parantez dışında sayıy yazsaydık."
	fmt.Println(veritipiniKendiBelirleyenYontem, reflect.TypeOf(veritipiniKendiBelirleyenYontem))

	//four metodh
	varsızDegisken := ":= var yazmak istemiyorsan bu işareti koy"
	fmt.Println(varsızDegisken)

	var intValue int8 = -123
	fmt.Println(intValue)

	//uint  : sadece positif değerler.
	//byte : pozitif çok küçük değerler

}
