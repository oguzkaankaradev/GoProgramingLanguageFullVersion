package main

import "fmt"

func main() {
	// java c# python gibi birçok dilde kullanılmaktadır.
	// golang dede interface vardır. oldukça fazla kullanacağız.
	// oyuzden siz ekran başındaki arkadaşlar kulaklarınızı açın dinleyin soruları en sonda son.

	//interfaceler imzama amacıyla kullanılır. amacı yazacağımız projenin özetlemektir.

	//var book1 = Book{desi: 10}
	// book := &Book{desi: 10}
	// book2 := &Book{desi: 20}
	// book3 := &Book{desi: 5}

	// fmt.Println("kARGO ÜCRETİNİZ", book.shipingCost())
	// fmt.Println("kARGO ÜCRETİNİZ", book2.shipingCost())
	// fmt.Println("kARGO ÜCRETİNİZ", book3.shipingCost())

	// books := []Book{{desi: 10}, {desi: 8}, {desi: 5}}

	//electronics := &Electronic{desi: 10}

	// electronics := []Electronic{{desi: 10}, {desi: 8}, {desi: 5}}
	// users := []string{"ali", "ahmet", "mehmet", "mustafa", "rıza"}
	// prices := []int{976456784567865678, 213, 343, 456, 254, 34678, 9876, 345454, 567, 78989, 232354, 9808987, 2343234, 78989}

	// fmt.Println(calculateTotalShipingCostBook(books))
	// fmt.Println("Elektronik kargo maliyeti:", calculateTotalShipingCostElectronics(electronics))

	var products []IShipibleProduct = []IShipibleProduct{
		&Book{desi: 10},
		&Book{desi: 12},
		&Book{desi: 15},
		&Electronic{desi: 12},
	}

	// var products []IShipibleProduct = []IShipibleProduct{
	// 	&Book{desi: 10},
	// 	&Book{desi: 20},
	// 	&Electronic{desi: 20},
	// }

	fmt.Println(calculateTotalShipingCost(products))

}

//Kargolama maliyetlerini hesaplayalım.
//elektronik, kitap veya bir gıdanın kargolama maliyeti farklı olabilir.

type IShipibleProduct interface {
	//contract tır. bir sçzleşmedir.

	shipingCost() int

	// var a string

}

// kitabın kargo maliyetini hesaplıyor
func (book *Book) shipingCost() int {
	return 5 + book.desi*2
}

// elektroniğin kargo maliyetini hesaplıyor
func (electronic *Electronic) shipingCost() int {
	return 10 + electronic.desi*3 //40
}

// kitap kargo maliyeti hesaplama
type Book struct {
	//desi = alan(bir kargonun ağırlığı, hacim(hacim demek kapladığı alandemel)), yani ne kadar alan kaplamaktadır.
	desi int //10
}

// elektronik kargo maliyeti
type Electronic struct {
	desi int
	//Siz burayi artırın, basit olması adına sadece desi üzerinde hesapluyoruz.
}

func calculateTotalShipingCost(products []IShipibleProduct) int {
	total := 0

	for _, product := range products {
		total += product.shipingCost()
	}

	return total
}

// Biz bir alışveriş sepeti yaptığımızı düşünelim.
// Toplam kargolama maliyetini hesaplayalım
// func calculateTotalShipingCostBook(books []Book) int {

// 	total := 0

// 	for _, book := range books {
// 		total += book.shipingCost()
// 	}

// 	return total //return içerisinde tutar, yazamaz nezamanki print edersek yazar.
// }

// func calculateTotalShipingCostElectronics(electronics []Electronic) int {
// 	total := 0

// 	for _, electronic := range electronics {
// 		total += electronic.shipingCost()
// 	}

// 	return total //return içerisinde tutar, yazamaz nezamanki print edersek yazar.
// }
