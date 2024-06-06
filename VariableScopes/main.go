package main

import "fmt"

//variable scope = değişkenlerin tanımlandığı yerler dışında kullanılmasıdır.
//4 farklı scopemiz vardır.
//* blok scpoe
//* fuction scope
//* global scope
//* package scpope = "ilerde öğrenecez"

//hiçbir bloğun veya fonksiyonun içinde olmayan scoptur.
//bu scopa o sayfadan her yerden ulaşabiliriz. fakat sayfa dışından ulaşamayız.
var globalScope = 10

func main() {
	var blockScope = true

	//block scope
	if blockScope {
		//SÜSLÜparantezlere blok denir.
		//bloğun içerisinde kullandığımız değişkenleri blok sışında kullanamadık.
		//bloğun içerisinde tanımlanan değişken sadece blok içerisinde kullanılır.
		var blockScope = 10
		fmt.Println(blockScope)
	}

	//bloğun dışında a değişkenini kullanamadık.
	//fmt.Println(a)

	fmt.Println(blockScope)

	//func scope
	test()
	// fmt.Println(funcScope)

	fmt.Println(globalScope)

}

func test() {

	//fonksiyon içersinde tanımladığımız değişkenlere başka bir fonksiyondan ulaşamayız.
	var funcScope = 10

	fmt.Println(funcScope)
}
