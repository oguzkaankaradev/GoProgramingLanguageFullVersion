package main

import "fmt"

// fonksiyonlar tüm programlama dillerinde olan yapılardır.
// Amacımız kodları gruplamaktır. Fonksiyonlarla yazmış olduğumuz kodları tekrar tekrar kullanabilmekteyiz.
//bir fonksiyon tanımlandıktan sonra gerek olması halinde programın her tarafında kullanılabilir.
//main fonksiyonunun asıl görevi kodları çalıştırmaktır. Örnek print fonksiyonu ekrana çıktıyı verir.

//önemli not!!!!!!!!! print ln yani birinci  yöntemi biz mainde çağırdığımızda direkt ekrana yazdırır.
//fakat ikinci yöntemi yani return ile çağırdığımızda biz ne zaman istersek o zaman ekrana yazdırır.
//Goda bir fonksiyon birden fazla değer dönebilir. değerden kasıt add fonksiyonunda parantezden sonraki int tir.
// eğer kullanmayacağımız bir tip varsa _(alttire) yazacağız
// ... ne kadar değer göndereceğimizi belli değilse kullanıyoruz.
func main() {
	//fmt.Print()

	//a ve b değerlerini müşteri girsin.
	//sonsuz bir toplama fonksiyonu yapın.

	//birinci yöntem çağırma
	//add(34567890, 567890)

	//ikinci yöntem çağırma
	// total := add(34567890, 567890)

	// fmt.Println(total)

	// total, difference, multiple := calculation(234234, 34234)

	// fmt.Println(total)
	// fmt.Println(difference)
	// fmt.Println(multiple)

	// var numbers = []int{1, 2, 3, 4, 5, 6, 6543, 678, 987654, 345678}
	// fmt.Println(sum(numbers))

	// fmt.Println(sum(3, 4, 5))

	fmt.Println(sum(3, 4, 5, 3456789, 234567890, 34567890, 34567890, 456789))
}

//basit bir fonksiyon tanımlayalım,
//fonksiyonun amacı gönderilen iki değeri toplamak olsun.

// birinci yöntem
// func add(x int, y int){
// 		fmt.Println(x + y)

// }

//ikinci yöntem
// func add(x int, y int) int {
// 	//func un içerisindeki değerleri topla ekrana yazır.

// 	return x + y
// }

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

//üçüncü yöntem
//Goda bir fonksiyon birden fazla değer dönebilir. değerden kasıt add fonksiyonunda parantezden sonraki int tir.

// func calculation(a int, b int) (int, int, int) {
// 	return a + b, a - b, a * b
// }

//4. kullanım arraylerle yaptık.
// func sum(numbers []int) int {
// 	sum := 0
// 	for _, value := range numbers {
// 		sum += value

// 	}
// 	return sum
// }

//
// func sum(x int, y int, z int) int {
// 	return x + y + z
// }

func sum(numbers ...int) int {

	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
