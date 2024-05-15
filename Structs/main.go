package main

import "fmt"

func main() {
	//Genelde banka, e ticaret uygulamaları gibi komplex uygulamaları parçalara bölerek yazarız, bu parçalar struct yardımı ile yapmaktayız.
	//örneğin bir sipariş kendi içerisinde birden fazla küçük yapılar olabilir herbir yapı için ayrı bir struct kullanmaktayız.
	//c# ve javada classlar vardırher bir classta bir veya birden fazla method vardır. Biz golangde bu methodlara struc deriz.
	//Struct içerisinde struct tanımlayabiliriz.
	//Biz c gibi dillerde structların içerisine bir fonksiyon tanımlayamıyorduk ancak golangde structların içerisine fonksiyon tanımlarız.
	//Structlar değer tiplidir.

	//veri tabanı olarak düşünün, ilk5 müşterinin geldiğini düşünün.
	var customer1 = Customer{id: 1, name: "oraz murat", lastName: "türkmenistan", age: 12, address: Address{city: "istanbul", street: "kavakpınar"}}

	//bir sutructdaki değişkenin değerini değiştirmek için bu şekilde kullanacağız.
	customer1.age = 18

	// fmt.Println(customer1)
	// customer1.ToStringFullName()

	// changeName(&customer1)
	// customer1.ToStringFullName()
	customer1.changeName("Türkmen akşabat")
	customer1.ToStringFullName()
}

type Customer struct {
	id       int32
	name     string
	lastName string
	age      int
	address  Address //structları struct içerisinde kullandık.

	//5000 tane bilgisi var.
}

type Address struct {
	city   string
	street string
}

//aşağıdaki örnekte fonksiyona struct verisi gönderilme örneği yapılmıştır. bir altında o prensibi ile daha uygun bir kullanımı verilmiştir.
//isim değiştirme fonksiyonu
// func changeName(customer *Customer) {
// 	customer.name = "muhammet murat"
// 	//int float string gibi veri tipleri ilkel veri tipidir. ilkel veri tipinde pointeri kullanmak için *customer koy
// 	//array slice struct gelişmiş veri tipidir. bunlarda customer yazmanız yeterli.
// }

//bestpractic
func (customer *Customer) changeName(newName string) {

	customer.name = newName
}

func (customer *Customer) ToStringFullName() {
	fmt.Printf("%s %s", customer.name, customer.lastName)
}
