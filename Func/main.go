package main

import "fmt"

// fonksiyonlar tüm programlama dillerinde olan yapılardır.
// Amacımız kodları gruplamaktır. Fonksiyonlarla yazmış olduğumuz kodları tekrar tekrar kullanabilmekteyiz.
//bir fonksiyon tanımlandıktan sonra gerek olması halinde programın her tarafında kullanılabilir.
//main fonksiyonunun asıl görevi kodları çalıştırmaktır. Örnek print fonksiyonu ekrana çıktıyı verir.

//önemli not!!!!!!!!! print ln yani birinci  yöntemi biz mainde çağırdığımızda direkt ekrana yazdırır.
//fakat ikinci yöntemi yani return ile çağırdığımızda biz ne zaman istersek o zaman ekrana yazdırır.
//Goda bir fonksiyon birden fazla değer dönebilir. değerden kasıt add fonksiyonunda parantezden sonraki int tir.

func main() {
	//fmt.Print()

	//a ve b değerlerini müşteri girsin.
	//sonsuz bir toplama fonksiyonu yapın.

	//birinci yöntem çağırma
	//add(34567890, 567890)

	//ikinci yöntem çağırma
	total := add(34567890, 567890)

	fmt.Println(total)

}

//basit bir fonksiyon tanımlayalım,
//fonksiyonun amacı gönderilen iki değeri toplamak olsun.

//birinci yöntem
// func add(x int, y int){
// 		fmt.Println(x + y)

// }

//ikinci yöntem
func add(x int, y int) int {
	//func un içerisindeki değerleri topla ekrana yazır.

	return x + y
}

/*
//üye girişi
//kart numarası
//password no
//fotoğraflar

func Crypt() {
	//burada şifreyi Gizlemek için kullanacağız.

	//artık şifre muhammet123. değilde
	// yeni şifre = akfçjaevnkxzçnafnwer35745tırowefsmdlvxdfbgwh5o
}
*/
