package main

import (
	"fmt"
	"sort"
)

func main() {

	//"sort" : karşılaştırma yapmak için kullanılır.
	//öncelikle soruyu yazalım.
	//Bir tamsayı dizisi verildiğinde, bu dizinin en büyük ve en küçük elemanlarının farkını bulan bir Go programı yazın.
	//ürünleriniz var filtreleme yapıyorsunuz, en yüksek fiyatı getir.

	// var kids [8]string

	// kids[0] = "Infant & Toddler Shoes"
	//array tanımlayalım - aşağıda arrayların farklı bir kullanım şekli mevcuttur.
	prices := []int{976456784567865678, 213, 343, 456, 254, 34678, 9876, 345454, 567, 78989, 232354, 9808987, 2343234, 78989}

	//en küçük ve en büyük değeri bulmak için kullanılır
	//küçükten büyüğe doğru sıralama işlemini yapmaktadır.
	sort.Ints(prices)

	//en küçük değeri bulalım
	minValue := prices[0]

	//en büyük fiyat yani yukarıdaki sırallamada en sondaki değeri alır.
	maxValue := prices[len(prices)-1]

	//farkı hesapla
	diffKarıHesapla := maxValue - minValue

	//en yüksek ürün ile en düşük ürünü al.
	//alımYapFiyatHesapla := maxValue + minValue

	//küçükten büyüğe sıralanışı
	fmt.Println(prices)

	fmt.Println("en düşük fiyat :", minValue, "en yüksek fiyat : ", maxValue, "Fiyatları karşılaştır karı hesapla", diffKarıHesapla)

}
